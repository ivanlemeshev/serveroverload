package ratelimiter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ivanlemeshev/serveroverload/internal/ratelimiter"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowCounter(t *testing.T) {
	rl := ratelimiter.NewFixedWindowCounter(2, 40*time.Millisecond)

	allowed := 0
	dropped := 0

	// Simulate 100 requests and count the number of allowed and dropped requests.
	for range 100 {
		if rl.IsAllowed() {
			allowed++
		} else {
			dropped++
		}
		// time.Sleep() is used to simulate the time it takes to process a request.
		time.Sleep(10 * time.Millisecond)
	}

	// The number of allowed and dropped requests should be close to each other
	// because the rate limiter allows 2 requests per 40ms window.
	// time.Sleep() does not guarantee the number of requests allowed or dropped.
	fmt.Println("allowed:", allowed, "|", "dropped:", dropped)
	assert.GreaterOrEqual(t, allowed, 49) // 49 or 51 requests should be allowed
	assert.GreaterOrEqual(t, dropped, 49) // 49 or 51 requests should be dropped
}
