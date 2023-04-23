package cache

import (
	"context"
	"github.com/cespare/xxhash"
	"github.com/pkg/errors"
	"sync"
	"time"
)

var _ Cache[int64] = (*cache[int64])(nil)

type Cache[T any] interface {
	Get(ctx context.Context, key string, fn func(context.Context) (*T, error)) (*T, error)
}

// bucketElement represent cached value with metadata
type bucketElement[T any] struct {
	value     *T
	key       string
	expiredAt int64
	next      *bucketElement[T]
	prev      *bucketElement[T]
}

// cacheBucket represent shard of cache.
// head, tail - part of linked list for LRU algorithm
type cacheBucket[T any] struct {
	sync.RWMutex
	data map[string]*bucketElement[T]
	head *bucketElement[T]
	tail *bucketElement[T]
}

type cache[T any] struct {
	sync.Mutex
	buckets    []*cacheBucket[T]
	config     Config
	bucketSize uint64
}

type Config struct {
	size        uint64 // total elements in cache
	bucketCount uint64
	ttl         time.Duration
}

func New[T any](config Config) Cache[T] {
	return &cache[T]{
		buckets:    make([]*cacheBucket[T], config.bucketCount),
		config:     config,
		bucketSize: config.size / config.bucketCount,
	}
}

// Get return cached T by key. If key not found or TTL expired - call fn and save returned value in cache.
func (c *cache[T]) Get(ctx context.Context, key string, fn func(context.Context) (*T, error)) (*T, error) {
	c.Lock()
	bucketNum := xxhash.Sum64String(key) % c.config.bucketCount
	bucket := c.buckets[bucketNum]
	if bucket == nil { // allocate bucket on first access
		bucket = &cacheBucket[T]{
			RWMutex: sync.RWMutex{},
			data:    make(map[string]*bucketElement[T], c.bucketSize),
			head:    nil,
			tail:    nil,
		}
		c.buckets[bucketNum] = bucket
	}
	c.Unlock()

	bucket.Lock()
	now := time.Now().Unix()
	element, ok := bucket.data[key]

	var value *T
	if !ok || element.expiredAt <= now { // request actual value if key not found or value expired
		bucket.Unlock() // allow edit bucket while request actual value via fn()
		var err error
		value, err = fn(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "call fn in cache Get")
		}
		bucket.Lock()
	}

	if !ok {
		// check is bucket has free space
		if uint64(len(bucket.data)) > c.bucketSize {
			bucket.removeOne(now)
		}

		element = &bucketElement[T]{
			key:       key,
			value:     value,
			expiredAt: now + int64(c.config.ttl),
		}

		// add new element as head
		bucket.head.next = element
		element.prev = bucket.head
		bucket.head = element
		bucket.data[key] = element
	} else if element != bucket.head {
		// detach from curr position
		element.next.prev = element.prev
		if element.prev != nil {
			element.prev.next = element.next
		} else {
			bucket.tail = element.next
		}

		// attach as head
		bucket.head.next = element
		element.prev = bucket.head
		element.next = nil
		bucket.head = element
	}

	if ok && element.expiredAt <= now {
		element.value = value
		element.expiredAt = now + int64(c.config.ttl)
	}

	bucket.Unlock()
	return element.value, nil
}

// removeOne remove one expired LRU element.
// If expired element not found - remove first LRU.
// Assume that mutex RLock called.
func (b *cacheBucket[T]) removeOne(now int64) {
	curr := b.tail
	for curr != nil && curr.expiredAt > now {
		curr = curr.next
	}

	if curr == nil { // expired element not found, remove LRU
		lruKey := b.tail.key
		b.tail = b.tail.next
		b.tail.prev = nil
		delete(b.data, lruKey)
		return
	}

	// remove expired
	curr.next.prev = curr.prev
	curr.prev.next = curr.next
	delete(b.data, curr.key)
}
