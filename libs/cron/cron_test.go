package cron

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func newTaskDesc(flag *atomic.Bool, err error, errCB ErrCallBack) TaskDescriptor {
	return TaskDescriptor{
		period: time.Second,
		task: func(ctx context.Context) error {
			flag.Store(true)
			return err
		},
		errCB:       errCB,
		retryPolicy: ImmediatelyAfterError,
	}
}

func TestCron(t *testing.T) {
	t.Run("should periodically run task", func(t *testing.T) {
		c := New()
		firstCallFlag := &atomic.Bool{}
		secondCallFlag := &atomic.Bool{}
		c.Add(gofakeit.MinecraftFood(), newTaskDesc(firstCallFlag, nil, nil))
		c.Add(gofakeit.MinecraftFood(), newTaskDesc(secondCallFlag, nil, nil))
		c.Stop()
		require.True(t, firstCallFlag.Load())
		require.True(t, secondCallFlag.Load())
	})

	t.Run("should periodically run task and report about err", func(t *testing.T) {
		c := New()
		firstCallFlag := &atomic.Bool{}
		secondCallFlag := &atomic.Bool{}
		secondCallFlagErr := &atomic.Bool{}
		c.Add(gofakeit.MinecraftFood(), newTaskDesc(firstCallFlag, nil, nil))
		c.Add(gofakeit.MinecraftFood(), newTaskDesc(secondCallFlag, errors.New(gofakeit.MinecraftBiome()), func(err error) {
			require.Error(t, err)
			secondCallFlagErr.Store(true)
		}))
		c.Stop()
		require.True(t, firstCallFlag.Load())
		require.True(t, secondCallFlag.Load())
		require.True(t, secondCallFlagErr.Load())
	})

}
