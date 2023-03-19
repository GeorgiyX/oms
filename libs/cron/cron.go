package cron

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type RetryPolicy int16

const (
	ImmediatelyAfterError RetryPolicy = iota
	ByScheduleAfterError
)

type Task func(ctx context.Context) error
type ErrCallBack func(err error)

type Cron interface {
	Add(name string, descriptor TaskDescriptor)
	Close()
}

type TaskDescriptor struct {
	period      time.Duration
	task        Task
	errCB       ErrCallBack
	retryPolicy RetryPolicy
}

type cron struct {
	wg         sync.WaitGroup
	stop       chan struct{}
	shouldStop atomic.Bool
}

func New() *cron {
	return &cron{
		wg:         sync.WaitGroup{},
		stop:       make(chan struct{}),
		shouldStop: atomic.Bool{},
	}
}

func (c *cron) Add(name string, descriptor TaskDescriptor) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		after := time.After(descriptor.period)
		for {
			var err error
			select {
			case <-after:
				err = descriptor.task(context.Background())
				if err != nil {
					descriptor.errCB(err)
				}
			}

			if c.shouldStop.Load() {
				return
			}

			if err != nil {
				after = afterErr(descriptor)
				continue
			}
			after = time.After(descriptor.period)
		}
	}()
}

func (c *cron) Stop() {
	c.shouldStop.Store(true)
	c.wg.Wait()
}

func afterErr(descriptor TaskDescriptor) <-chan time.Time {
	switch descriptor.retryPolicy {
	case ImmediatelyAfterError:
		return time.After(time.Millisecond * 100)
	case ByScheduleAfterError:
		return time.After(descriptor.period)
	default:
		return time.After(descriptor.period)
	}
}
