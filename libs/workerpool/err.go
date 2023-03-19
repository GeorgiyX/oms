package workerpool

import (
	"sync"

	"go.uber.org/multierr"
)

// errHandler help to stop pool if err happen
type errHandler struct {
	err      error
	stopOnce sync.Once // we don't want panic on double close
	sync.Mutex
	errHappen chan struct{}
}

func newErrHandler() *errHandler {
	return &errHandler{
		err:       nil,
		stopOnce:  sync.Once{},
		Mutex:     sync.Mutex{},
		errHappen: make(chan struct{}),
	}
}

// Error implement error interface
func (e *errHandler) Error() error {
	e.Lock()
	defer e.Unlock()
	return e.err
}

func (e *errHandler) Close() {
	e.stopOnce.Do(func() { close(e.errHappen) })
}

// Happen return chan to check is error happen
func (e *errHandler) Happen() chan struct{} {
	return e.errHappen
}

// Set error and close signal chan
func (e *errHandler) Set(err error) {
	e.stopOnce.Do(func() { close(e.errHappen) }) // signal that err happen
	e.Lock()
	defer e.Unlock()
	e.err = multierr.Append(e.err, err) // save error
}
