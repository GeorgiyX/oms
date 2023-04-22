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

type bucketElement[T any] struct {
	value     *T
	key       string
	expiredAt int64
	next      *bucketElement[T]
	prev      *bucketElement[T]
}

type cacheBucket[T any] struct {
	sync.RWMutex
	data map[string]*bucketElement[T]
	head *bucketElement[T]
	tail *bucketElement[T]
}

type cache[T any] struct {
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
		buckets:    make([]*cacheBucket[T], 0, config.bucketCount),
		config:     config,
		bucketSize: config.size / config.bucketCount,
	}
}

// Get return cachet T by key. If key not found or TTL expired - call fn and save returned value in cache.
func (c *cache[T]) Get(ctx context.Context, key string, fn func(context.Context) (*T, error)) (*T, error) {
	bucketNum := xxhash.Sum64String(key) % c.config.bucketCount
	bucket := c.buckets[bucketNum]

	bucket.Lock()
	if bucket.data == nil {
		bucket.data = make(map[string]*bucketElement[T], c.bucketSize)
	}

	now := time.Now().Unix()
	element, ok := bucket.data[key]

	var value *T
	if !ok || element.expiredAt <= now {
		bucket.Unlock() // allow edit bucket while request actual value via fn()
		var err error
		value, err = fn(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "call fn in cache Get")
		}

		element.value = value
		element.expiredAt = now + int64(c.config.ttl)
	}

	bucket.Lock()
	if !ok {
		element = &bucketElement[T]{
			key:       key,
			value:     value,
			expiredAt: now,
		}

		// add new element as head
		bucket.head.next = element
		element.prev = bucket.head
		bucket.head = element

		// check is bucket has free space
		if uint64(len(bucket.data)) > c.bucketSize {
			bucket.removeOne(now)
		}
	} else if element != bucket.head {
		// detach from curr position
		element.next.prev = element.prev
		if element.prev != nil {
			element.prev.next = element.next
		}

		// attach as head
		bucket.head.next = element
		element.prev = bucket.head
		element.next = nil
		element.expiredAt = now
		bucket.head = element
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
