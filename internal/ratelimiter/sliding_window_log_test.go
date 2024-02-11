package ratelimiter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ivanlemeshev/serveroverload/internal/ratelimiter"
)

func TestSlidingWindowLog(t *testing.T) {
	rl := ratelimiter.NewSlidingWindowLog(2, 40*time.Millisecond)

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

	// time.Sleep() does not guarantee the number of requests allowed or dropped.
	fmt.Println("allowed:", allowed, "|", "dropped:", dropped)
	assert.GreaterOrEqual(t, allowed, 50)
	assert.GreaterOrEqual(t, dropped, 50)
}
