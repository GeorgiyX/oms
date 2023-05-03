package cache

import (
	"context"
	"testing"
	"time"
	"unsafe"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

const (
	bucketsCount = 32
)

func getValueFuncBench() GetFunc[int64] {
	return func(ctx context.Context) (*int64, error) {
		return &someNum, nil
	}
}

func getSizeOfBucketElement() int64 {
	return int64(unsafe.Sizeof(bucketElement[int64]{}))
}

type keysStore []string

func (k *keysStore) random() string {
	return (*k)[gofakeit.IntRange(0, len(*k)-1)]
}

func genKeys(count uint64) keysStore {
	keys := make([]string, 0, count)
	for i := uint64(0); i < count; i++ {
		key := gofakeit.BitcoinAddress()
		keys = append(keys, key)
	}
	return keys
}

func prepareBenchCache(b *testing.B, cacheSize, keysCount uint64, ttl time.Duration) (Cache[int64], keysStore) {
	c, err := New[int64](Config{
		Size:        cacheSize,
		BucketCount: bucketsCount,
		TTL:         ttl,
		Name:        gofakeit.BeerName(),
	})
	require.NoError(b, err)

	ctx := context.Background()
	keys := genKeys(keysCount)
	for _, key := range keys {
		_, err := c.Get(ctx, key, getValueFuncBench())
		require.NoError(b, err)
	}

	return c, keys
}

func BenchmarkCache(b *testing.B) {
	ctx := context.Background()
	b.Run("cache hit read (100% hit)", func(b *testing.B) {
		c, keys := prepareBenchCache(b, 20_000, 18_000, time.Hour)
		b.SetBytes(getSizeOfBucketElement())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = c.Get(ctx, keys.random(), getValueFuncBench())
		}
	})

	b.Run("cache miss read (100% miss = 50% overfull + 50% ttl)", func(b *testing.B) {
		c, keys := prepareBenchCache(b, 20_000, 40_000, -time.Hour)
		b.SetBytes(getSizeOfBucketElement())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = c.Get(ctx, keys.random(), getValueFuncBench())
		}
	})

	b.Run("cache mix read (25% miss (overfull), 75% hit)", func(b *testing.B) {
		c, keys := prepareBenchCache(b, 20_000, 25_000, time.Hour)
		b.SetBytes(getSizeOfBucketElement())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = c.Get(ctx, keys.random(), getValueFuncBench())
		}
	})
}
