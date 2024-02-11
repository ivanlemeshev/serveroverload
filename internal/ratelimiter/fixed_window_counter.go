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
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Since(rl.windowStartTime) > rl.windowSize {
		rl.requestCount = 0
		rl.windowStartTime = time.Now()
	}

	if rl.requestCount < rl.limit {
		rl.requestCount++
		return true
	}

	return false
}
