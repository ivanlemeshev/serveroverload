package ratelimiter

import (
	"sync"
	"time"
)

// SlidingWindowCounter is a rate limiter that uses a sliding window counter
// algorithm.
type SlidingWindowCounter struct {
	mu                         sync.Mutex
	previousWindowRequestCount float64 // number of requests in the previous window
	currentWindowRequestCount  float64 // number of requests in the current window
	limit                      int     // max number of requests in the window
	windowSize                 time.Duration
	windowStartTime            time.Time
}

// NewSlidingWindowCounter returns a new SlidingWindowCounter with the given
// limit and window size.
func NewSlidingWindowCounter(limit int, windowSize time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		previousWindowRequestCount: 0,
		currentWindowRequestCount:  0,
		limit:                      limit,
		windowSize:                 windowSize,
		windowStartTime:            time.Now(),
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (rl *SlidingWindowCounter) IsAllowed() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Since(rl.windowStartTime) > rl.windowSize {
		rl.previousWindowRequestCount = rl.currentWindowRequestCount
		rl.currentWindowRequestCount = 0
		rl.windowStartTime = time.Now()
	}

	previousWindowTimeFrame := (rl.windowSize.Seconds() - time.Since(rl.windowStartTime).Seconds()) / rl.windowSize.Seconds()
	weightedRequestCount := rl.previousWindowRequestCount*previousWindowTimeFrame + rl.currentWindowRequestCount

	if int(weightedRequestCount) < rl.limit {
		rl.currentWindowRequestCount++
		return true
	}

	return false
}
