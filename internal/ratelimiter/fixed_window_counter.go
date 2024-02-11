package ratelimiter

import (
	"sync"
	"time"
)

// FixedWindowCounter is a rate limiter that uses a fixed window counter
// algorithm.
type FixedWindowCounter struct {
	mu              sync.Mutex
	requestCount    int           // number of requests in the current window
	limit           int           // max number of requests allowed in the window
	windowSize      time.Duration // time window for the rate limit
	windowStartTime time.Time     // start of the current window
}

// NewFixedWindowCounter returns a new fixed window counter rate limiter with a
// given request limit and window size.
func NewFixedWindowCounter(limit int, windowSize time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		requestCount:    0,
		limit:           limit,
		windowSize:      windowSize,
		windowStartTime: time.Now(),
	}
}

// IsAllowed returns true if the rate limiter allows the request, and false if
// it does not.
func (rl *FixedWindowCounter) IsAllowed() bool {
	// Lock the rate limiter while we update its state.
	// This is necessary because the requestCount and windowStartTime fields are
	// shared between multiple goroutines.
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// If the current time is greater than the start of the next window, reset
	// the request count and start time.
	if time.Since(rl.windowStartTime) > rl.windowSize {
		rl.requestCount = 0
		rl.windowStartTime = time.Now()
	}

	// If the number of requests in the window is less than the limit, increment
	// the request count and return true.
	if rl.requestCount < rl.limit {
		rl.requestCount++
		return true
	}

	// If the number of requests in the window is equal or greater than the
	// limit, return false.
	return false
}
