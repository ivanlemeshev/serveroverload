package ratelimiter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ivanlemeshev/serveroverload/internal/ratelimiter"
)

func TestTokenBucket(t *testing.T) {
	rl := ratelimiter.NewTokenBucket(10, 10)

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
		time.Sleep(40 * time.Millisecond)
	}

	// The number of allowed and dropped requests should be close to each other.
	// time.Sleep() does not guarantee the number of requests allowed or dropped.
	fmt.Println("allowed:", allowed, "|", "dropped:", dropped)
	assert.GreaterOrEqual(t, allowed, 49) // 49 or 51 requests should be allowed
	assert.GreaterOrEqual(t, dropped, 49) // 49 or 51 requests should be dropped
}
