package workerpool

import (
	"context"
	"sync"
)

type TaskFn func(ctx context.Context) error
type taskFnWrapper func() error

type WorkerPool interface {
	Schedule(ctx context.Context, task TaskFn)
	Stop() error
}

var _ WorkerPool = (*pool)(nil)

type pool struct {
	tasks    chan taskFnWrapper
	stopOnce sync.Once // we don't want panic on double close
	wg       sync.WaitGroup
	err      *errHandler
}

func New(workersCount, queueCapacity int) *pool {
	p := &pool{
		tasks:    make(chan taskFnWrapper, queueCapacity),
		stopOnce: sync.Once{},
		wg:       sync.WaitGroup{},
		err:      newErrHandler(),
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
		for {
			// first check error in pool
			select { // check one of: err or new task
			case <-p.err.Happen():
				return // pool stop if err happen in any worker
			default:
			}

			// if no error - check for task
			select {
			case task, ok := <-p.tasks:
				if !ok { // task channel closed, all tasks done
					return
				}
				err := task() // do task
				if err != nil {
					p.err.Set(err) // report err
				}
			default:
			}
		}
	}()

}

// Schedule add new task to worker pool.
// This call may block if added more tasks than the work pool chan buffer size
func (p *pool) Schedule(ctx context.Context, task TaskFn) {
	p.tasks <- func() error { // add task to pool
		select {
		case <-ctx.Done(): // don't do task with ended context
			return nil
		default:
			return task(ctx)
		}
	}
}

// Stop close task channel and wait workers stopping. All Tasks will be processed.
// Do not call Schedule after Stop call, also as in another goroutine.
func (p *pool) Stop() error {
	p.stopOnce.Do(func() {
		close(p.tasks) // signal workers that pool going stop
		p.wg.Wait()    // wait goroutines exit.
		p.err.Close()
	})
	return p.err.Error()
}
