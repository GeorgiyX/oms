package ratelimiter

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrWillExceedDeadLine = errors.New("ratelimiter: wait with this ctx would exceed context deadline")
)

// Config represent requests (aka token) per interval (aka buket length)
type Config struct {
	interval time.Duration // buket length
	requests uint64        // tokens in bucket
}

type ratelimiter struct {
	sync.Mutex
	end    time.Time // moment when bucket end
	tokens uint64    // remain tokens in buket
	config Config
}

func New(config Config) *ratelimiter {
	return &ratelimiter{
		Mutex:  sync.Mutex{},
		end:    time.Now().Add(config.interval),
		tokens: 0,
		config: config,
	}
}

func (r *ratelimiter) Wait(ctx context.Context) error {
	r.Lock()
	now := time.Now()
	r.updateBucket(now)           // check and update if bucket don't have enough token or too old
	waitTo := r.bucketStartTime() // moment until we must wait

	if r.isAfterDeadline(ctx, waitTo) {
		r.Unlock()
		return ErrWillExceedDeadLine
	}

	r.takeToken() // reserve place in last buket
	r.Unlock()

	after := time.After(waitTo.Sub(now))

	select {
	case <-after:
		return nil // time to do request
	case <-ctx.Done():
		return ctx.Err() // ctx executed faster
	}
}

func (r *ratelimiter) updateBucket(now time.Time) {
	if r.end.After(now) && r.tokens != 0 { // update only if bucket in past or tokens end
		return
	}

	r.tokens = r.config.requests

	if r.end.After(now) {
		r.end = r.end.Add(r.config.interval) // reserve next bucket
		return
	}

	r.end = now.Add(r.config.interval) // bucket start right now
}

func (r *ratelimiter) isAfterDeadline(ctx context.Context, moment time.Time) bool {
	deadline, ok := ctx.Deadline()
	if !ok {
		return false
	}
	return deadline.Before(moment)
}

func (r *ratelimiter) bucketStartTime() time.Time {
	return r.end.Add(-1 * r.config.interval)
}

func (r *ratelimiter) takeToken() {
	r.tokens--
}
