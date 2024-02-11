package ratelimiter

import (
	"math"
	"sync"
	"time"
)

// TokenBucket is a rate limiter that uses a token bucket algorithm.
type TokenBucket struct {
	mu             sync.Mutex
	currentTokens  float64   // current number of tokens in the bucket
	bucketSize     float64   // max number of tokens in the bucket
	refillRate     float64   // number of tokens to add per second
	lastRefillTime time.Time // last time the bucket was refilled
}

// NewTokenBucket creates a new TokenBucket with the
// given bucket size and refill rate.
func NewTokenBucket(bucketSize, refillRate float64) *TokenBucket {
	return &TokenBucket{
		currentTokens:  bucketSize,
		bucketSize:     bucketSize,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (rl *TokenBucket) IsAllowed() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refill()

	if int(rl.currentTokens) > 0 {
		rl.currentTokens--
		return true
	}

	return false
}

// refill adds tokens to the bucket based on the refill rate.
func (rl *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(rl.lastRefillTime)
	tokensToAdd := rl.refillRate * duration.Seconds()
	rl.currentTokens = math.Min(rl.bucketSize, rl.currentTokens+tokensToAdd)
	rl.lastRefillTime = now
}
