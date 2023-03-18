package workerpool

import (
	"context"
	"sync"
)

type TaskFn func(ctx context.Context)
type taskFnWrapper func()

type WorkerPool interface {
	Schedule(ctx context.Context, task TaskFn)
	Stop()
}

var _ WorkerPool = (*pool)(nil)

type pool struct {
	tasks    chan taskFnWrapper
	stopOnce sync.Once // We don't want to close the channel again
	wg       sync.WaitGroup
}

func New(workersCount, queueCapacity int) *pool {
	p := &pool{
		tasks: make(chan taskFnWrapper, workersCount),
		wg:    sync.WaitGroup{},
	}

	for i := 0; i < workersCount; i++ {
		p.runWorker()
	}

	return p
}

// runWorker start worker goroutine
func (p *pool) runWorker() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for task := range p.tasks { // process tasks
			task()
		}
	}()

}

// Schedule add new task to worker pool.
// This call may block if added more tasks than the work pool chan buffer size
func (p *pool) Schedule(ctx context.Context, task TaskFn) {
	p.tasks <- func() { // Add task to pool
		select {
		case <-ctx.Done(): // Don't do task with ended context
			return
		default:
			task(ctx)
		}
	}
}

// Stop close task channel and what workers stopping. All Tasks will be processed.
// Do not call Schedule after Stop call, also as in another goroutine.
func (p *pool) Stop() {
	p.stopOnce.Do(func() {
		close(p.tasks) // Signal workers that pool going stop
		p.wg.Wait()    // Wait goroutines exit.
	})
}
