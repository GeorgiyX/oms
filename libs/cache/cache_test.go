package cache

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
	"time"
)

var someNum = int64(42)

func getValueFunc(isErr bool) (func(ctx context.Context) (*int64, error), *atomic.Int64) {
	counter := &atomic.Int64{}
	return func(ctx context.Context) (*int64, error) {
		counter.Add(1)
		if !isErr {
			return &someNum, nil
		}
		return &someNum, errors.New(gofakeit.MinecraftBiome())
	}, counter
}

func TestCache(t *testing.T) {
	ctx := context.Background()
	t.Run("should correct store elements", func(t *testing.T) {
		c, err := New[int64](Config{
			size:        10,
			bucketCount: 2,
			ttl:         time.Minute,
		})
		require.NoError(t, err)

		fn, counter := getValueFunc(false)
		key1 := gofakeit.BeerName()

		value, err := c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(1), counter.Load())

		value, err = c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(1), counter.Load())

		key2 := gofakeit.BeerName()
		value, err = c.Get(ctx, key2, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(2), counter.Load())
	})

	t.Run("should replace LRU if cache full", func(t *testing.T) {
		c, err := New[int64](Config{
			size:        1,
			bucketCount: 1,
			ttl:         time.Minute,
		})
		require.NoError(t, err)

		fn, counter := getValueFunc(false)
		key1 := gofakeit.BeerName()

		value, err := c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(1), counter.Load())

		key2 := gofakeit.BeerName()
		value, err = c.Get(ctx, key2, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(2), counter.Load())

		value, err = c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(3), counter.Load())
	})

	t.Run("should refresh value if value expired", func(t *testing.T) {
		c, err := New[int64](Config{
			size:        10,
			bucketCount: 2,
			ttl:         0,
		})
		require.NoError(t, err)

		fn, counter := getValueFunc(false)
		key1 := gofakeit.BeerName()

		value, err := c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(1), counter.Load())

		time.Sleep(time.Millisecond)

		value, err = c.Get(ctx, key1, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(2), counter.Load())
	})

	t.Run("should return err if refresh value fail", func(t *testing.T) {
		c, err := New[int64](Config{
			size:        10,
			bucketCount: 2,
			ttl:         -time.Minute,
		})
		require.NoError(t, err)

		fnErr, counterErr := getValueFunc(true)
		key := gofakeit.BeerName()

		value, err := c.Get(ctx, key, fnErr)
		require.Error(t, err)
		require.Nil(t, value)
		require.Equal(t, int64(1), counterErr.Load())

		fn, counter := getValueFunc(false)

		value, err = c.Get(ctx, key, fn)
		require.NoError(t, err)
		require.NotNil(t, value)
		require.Equal(t, someNum, *value)
		require.Equal(t, int64(1), counter.Load())
	})
}
