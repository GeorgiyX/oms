package workerpool

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	tasks = 100
)

func TestWorkerPool(t *testing.T) {
	t.Run("should correctly process tasks (10 goroutines)", func(t *testing.T) {
		calls := atomic.Int64{}
		p := New(10, tasks)

		for i := 0; i < tasks; i++ {
			p.Schedule(context.Background(), func(ctx context.Context) error {
				calls.Add(1)
				return nil
			})
		}
		err := p.Stop()

		require.NoError(t, err)
		require.Equal(t, tasks, int(calls.Load()))
	})

	t.Run("should stop if err happen (1 goroutine)", func(t *testing.T) {
		expectNum := tasks / 2
		calls := atomic.Int64{}
		p := New(1, tasks)

		for i := 0; i < tasks; i++ {
			var err error
			if i == expectNum-1 {
				err = errors.New(gofakeit.MinecraftBiome())
			}

			p.Schedule(context.Background(), func(ctx context.Context) error {
				calls.Add(1)
				return err
			})
		}
		err := p.Stop()

		require.Error(t, err)
		require.Equal(t, expectNum, int(calls.Load()))
	})

	t.Run("should save error (10 goroutines)", func(t *testing.T) {
		expectNum := tasks / 2
		p := New(10, tasks)

		for i := 0; i < tasks; i++ {
			var err error
			if i == expectNum {
				err = errors.New(gofakeit.MinecraftBiome())
			}

			p.Schedule(context.Background(), func(ctx context.Context) error {
				return err
			})
		}
		err := p.Stop()

		require.Error(t, err)
	})
}
