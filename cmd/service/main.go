package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ivanlemeshev/serveroverload/internal/middleware"
	"github.com/ivanlemeshev/serveroverload/internal/overloaddetector"
	"github.com/ivanlemeshev/serveroverload/internal/ratelimiter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Simulate request processing time
	time.Sleep(20 * time.Millisecond)
	// Response with status code 200
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.HandleFunc("/", handler)

	fwcrl := ratelimiter.NewFixedWindowCounter(30, 1*time.Second)
	http.HandleFunc("/fixed_window_counter", middleware.RateLimiting(fwcrl, handler))

	tbrl := ratelimiter.NewTokenBucket(30, 30)
	http.HandleFunc("/token_bucket", middleware.RateLimiting(tbrl, handler))

	lbrl := ratelimiter.NewLeakyBucket(30, 100*time.Millisecond, 5)
	http.HandleFunc("/leaky_bucket", middleware.RateLimiting(lbrl, handler))

	swlrl := ratelimiter.NewSlidingWindowLog(30, 1*time.Second)
	http.HandleFunc("/sliding_window_log", middleware.RateLimiting(swlrl, handler))

	swcrl := ratelimiter.NewSlidingWindowCounter(30, 1*time.Second)
	http.HandleFunc("/sliding_window_counter", middleware.RateLimiting(swcrl, handler))

	od := overloaddetector.New(ctx, 20*time.Millisecond, 21*time.Millisecond)
	http.HandleFunc("/overload_detector", middleware.OverloadDetecting(od, handler))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
