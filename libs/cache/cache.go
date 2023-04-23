package cache

import (
	"context"
	"sync"
	"time"

	"github.com/cespare/xxhash"
	"github.com/pkg/errors"
)

var _ Cache[int64] = (*cache[int64])(nil)

type GetFunc[T any] func(context.Context) (*T, error)

type Cache[T any] interface {
	Get(ctx context.Context, key string, fn GetFunc[T]) (*T, error)
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
	Size        uint64 // total elements in cache
	BucketCount uint64
	TTL         time.Duration
	Name        string
}

func New[T any](config Config) (Cache[T], error) {
	if config.BucketCount == 0 {
		return nil, errors.New("zero bucket count in config")
	}

	if config.Size < config.BucketCount {
		return nil, errors.New("zero cache size count in config")
	}

	return &cache[T]{
		buckets:    make([]*cacheBucket[T], config.BucketCount),
		config:     config,
		bucketSize: config.Size / config.BucketCount,
	}, nil
}

// Get return cached T by key. If key not found or TTL expired - call fn and save returned value in cache.
func (c *cache[T]) Get(ctx context.Context, key string, fn GetFunc[T]) (*T, error) {
	requestCounter.WithLabelValues(c.config.Name).Inc()
	exitCountRtCache := countRtCache(c.config.Name)
	defer exitCountRtCache()

	c.Lock()
	bucketNum := xxhash.Sum64String(key) % c.config.BucketCount
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

		missCounter.WithLabelValues(c.config.Name).Inc()
		exitCountRtFn := countRtFn(c.config.Name)
		defer exitCountRtFn()

		var err error
		value, err = fn(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "call fn in cache Get")
		}

		bucket.Lock()
	}

	// set element position at LRU linked list
	// buket.head - most resent used element, buket.tail - least resent used
	if !ok {
		// check is bucket has free space
		if uint64(len(bucket.data)) >= c.bucketSize {
			err := bucket.removeOne(now)
			if err != nil {
				bucket.Unlock()
				return nil, err
			}
		}

		element = &bucketElement[T]{
			key:       key,
			value:     value,
			expiredAt: now + int64(c.config.TTL),
		}

		// add new element as head
		if bucket.head != nil {
			bucket.head.next = element
			element.prev = bucket.head
		}
		bucket.head = element
		bucket.data[key] = element

		if bucket.tail == nil {
			bucket.tail = element
		}
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
		element.expiredAt = now + int64(c.config.TTL)
	}

	bucket.Unlock()
	return element.value, nil
}

// removeOne remove one expired LRU element.
// If expired element not found - remove first LRU.
// Assume that mutex RLock called.
func (b *cacheBucket[T]) removeOne(now int64) error {
	if b.tail == nil || b.head == nil || len(b.data) == 0 {
		return errors.New("incorrect cache size configuration")
	}

	curr := b.tail
	for curr != nil && curr.expiredAt > now {
		curr = curr.next
	}

	if curr == nil { // expired element not found, remove LRU
		lruKey := b.tail.key

		if b.tail == b.head {
			b.head = nil
			b.tail = nil
		} else {
			b.tail = b.tail.next
			b.tail.prev = nil
		}

		delete(b.data, lruKey)
		return nil
	}

	// remove expired
	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}

	if curr == b.head {
		b.head = curr.prev
	}

	if curr == b.tail {
		b.tail = curr.next
	}
	delete(b.data, curr.key)
	return nil
}
