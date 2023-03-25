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

var _ Cron = (*cron)(nil)

type Cron interface {
	Add(descriptor TaskDescriptor)
	Stop()
}

type TaskDescriptor struct {
	Period      time.Duration
	Task        Task
	ErrCB       ErrCallBack
	RetryPolicy RetryPolicy
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

func (c *cron) Add(descriptor TaskDescriptor) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		after := time.After(descriptor.Period)
		for {
			var err error
			select {
			case <-after:
				err = descriptor.Task(context.Background())
				if err != nil {
					descriptor.ErrCB(err)
				}
			}

			if c.shouldStop.Load() {
				return
			}

			if err != nil {
				after = afterErr(descriptor)
				continue
			}
			after = time.After(descriptor.Period)
		}
	}()
}

func (c *cron) Stop() {
	c.shouldStop.Store(true)
	c.wg.Wait()
}

func afterErr(descriptor TaskDescriptor) <-chan time.Time {
	switch descriptor.RetryPolicy {
	case ImmediatelyAfterError:
		return time.After(time.Millisecond * 100)
	case ByScheduleAfterError:
		return time.After(descriptor.Period)
	default:
		return time.After(descriptor.Period)
	}
}
