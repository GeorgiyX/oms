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

// Config represent Requests (aka token) per Interval (aka buket length)
type Config struct {
	Interval time.Duration // buket length
	Requests uint64        // tokens in bucket
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
		end:    time.Now().Add(config.Interval),
		tokens: 0,
		config: config,
	}
}

// Wait is completed no more often than "Requests" once in the "Interval".
// The algorithm splits the timeline into buckets with a duration of "Interval"
// on which Wait method can terminate no more than "Requests" times.
// The algorithm selects nearest bucket on and waits until the moment of occurrence
// of the bucket. If the context deadline comes earlier, an error is returned.
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

	r.tokens = r.config.Requests

	if r.end.After(now) {
		r.end = r.end.Add(r.config.Interval) // reserve next bucket
		return
	}

	r.end = now.Add(r.config.Interval) // bucket start right now
}

func (r *ratelimiter) isAfterDeadline(ctx context.Context, moment time.Time) bool {
	deadline, ok := ctx.Deadline()
	if !ok {
		return false
	}
	return deadline.Before(moment)
}

func (r *ratelimiter) bucketStartTime() time.Time {
	return r.end.Add(-1 * r.config.Interval)
}

func (r *ratelimiter) takeToken() {
	r.tokens--
}
