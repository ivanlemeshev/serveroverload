package ratelimiter

import (
	"time"
)

// LeakyBucket is a rate limiter that uses the leaky bucket algorithm.
type LeakyBucket struct {
	bucket      chan struct{} // channel that acts as a queue for the leaky bucket
	outflowRate int           // number of requests to remove from the bucket per interval
	interval    time.Duration // interval at which to remove requests from the bucket
	lastLeak    time.Time     // last time the bucket was leaked
}

// NewLeakyBucket returns a new LeakyBucket with the given bucket size,
// interval, and outflow rate.
func NewLeakyBucket(bucketSize int, interval time.Duration, outflowRate int) *LeakyBucket {
	return &LeakyBucket{
		bucket:      make(chan struct{}, bucketSize),
		outflowRate: outflowRate,
		interval:    interval,
		lastLeak:    time.Now(),
	}
}

// IsAllowed returns true if the rate limiter allows the request, and false
// otherwise.
func (rl *LeakyBucket) IsAllowed() bool {
	rl.leak()

	// This is a non-blocking send operation. If the bucket is full, the
	// default case will be selected and the function will return false.
	select {
	case rl.bucket <- struct{}{}:
		return true
	default:
		return false
	}
}

// leak removes requests from the bucket at the outflow rate.
func (rl *LeakyBucket) leak() {
	if time.Since(rl.lastLeak) >= rl.interval {
		for range rl.outflowRate {
			select {
			case <-rl.bucket:
			default:
				break
			}
		}
		rl.lastLeak = time.Now()
	}
}
