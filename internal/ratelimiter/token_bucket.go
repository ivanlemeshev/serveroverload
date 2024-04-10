package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type bucket struct {
	currentTokens  float64
	lastRefillTime time.Time
}

type TokenBucket struct {
	mu         sync.Mutex
	buckets    map[string]*bucket
	bucketSize float64
	refillRate float64
}

func NewTokenBucket(bucketSize, refillRate float64) *TokenBucket {
	return &TokenBucket{
		bucketSize: bucketSize,
		refillRate: refillRate,
		buckets:    make(map[string]*bucket),
	}
}

func (rl *TokenBucket) IsAllowed(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if _, ok := rl.buckets[key]; !ok {
		rl.buckets[key] = &bucket{
			currentTokens:  rl.bucketSize - 1,
			lastRefillTime: time.Now(),
		}
		return true
	}
	rl.refill(key)
	if int(rl.buckets[key].currentTokens) > 0 {
		rl.buckets[key].currentTokens--
		return true
	}
	return false
}

func (rl *TokenBucket) refill(key string) {
	now := time.Now()
	duration := now.Sub(rl.buckets[key].lastRefillTime)
	tokensToAdd := rl.refillRate * duration.Seconds()
	rl.buckets[key].currentTokens = math.Min(rl.bucketSize,
		rl.buckets[key].currentTokens+tokensToAdd)
	rl.buckets[key].lastRefillTime = now
}
