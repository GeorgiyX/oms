package ratelimiter

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	t.Run("should block if limit", func(t *testing.T) {
		rl := New(Config{
			Interval: time.Second,
			Requests: 2,
		})

		start := time.Now()
		for i := 0; i < 10; i++ {
			err := rl.Wait(context.Background())
			require.NoError(t, err)
		}
		end := time.Now()
		require.GreaterOrEqual(t, end.Sub(start), time.Second*4)
	})

	t.Run("should return error if deadline happen before wait time end", func(t *testing.T) {
		rl := New(Config{
			Interval: time.Second * 3,
			Requests: 1,
		})

		ctx := context.Background()
		err := rl.Wait(ctx)
		require.NoError(t, err)

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
		defer cancel()
		err = rl.Wait(ctx)
		require.EqualError(t, err, ErrWillExceedDeadLine.Error())
	})
}
