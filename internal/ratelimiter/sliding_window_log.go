package ratelimiter

import (
	"sync"
	"time"
)

// SlidingWindowLog is a rate limiter that uses a sliding window log algorithm.
type SlidingWindowLog struct {
	mu         sync.Mutex
	limit      int           // max number of requests allowed in the window
	windowSize time.Duration // time window for the rate limit
	requestLog []int64       // log of request timestamps
}

// NewSlidingWindowLog returns a new sliding window log rate limiter with a
// given request limit and window size.
func NewSlidingWindowLog(limit int, windowSize time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		limit:      limit,
		windowSize: windowSize,
		requestLog: make([]int64, 0),
	}
}

// IsAllowed returns true if the rate limiter allows the request, and false if
// it does not.
func (rl *SlidingWindowLog) IsAllowed() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	requestTime := time.Now().UnixMicro()
	slidingWindowStartTime := requestTime - rl.windowSize.Microseconds()

	rl.removeOutdatedTimestamps(slidingWindowStartTime)

	if len(rl.requestLog) < rl.limit {
		rl.logTimestamp(requestTime)
		return true
	}

	return false
}

// logTimestamp adds a request timestamp to the request log.
func (rl *SlidingWindowLog) logTimestamp(requestTime int64) {
	rl.requestLog = append(rl.requestLog, requestTime)
}

// removeOutdatedTimestamps removes timestamps from the request log that are
// outside the sliding window.
func (rl *SlidingWindowLog) removeOutdatedTimestamps(windowStartTime int64) {
	found := false
	startWindowIndex := 0
	for i, t := range rl.requestLog {
		if t >= windowStartTime {
			found = true
			startWindowIndex = i
			break
		}
	}

	if found {
		rl.requestLog = rl.requestLog[startWindowIndex:]
	} else {
		rl.requestLog = rl.requestLog[:0]
	}
}
