package middleware

import (
	"fmt"
	"net/http"
)

// RateLimiter is an interface for rate limiting.
type RateLimiter interface {
	IsAllowed() bool
}

// RateLimiting is a middleware that limits the number of requests.
// It returns a 429 status code if the request is not allowed.
// Otherwise, it calls the original handler.
func RateLimiting(rl RateLimiter, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is allowed.
		if !rl.IsAllowed() {
			// Response with status code 429
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, http.StatusText(http.StatusTooManyRequests))
			return
		}
		// Call the original handler if the request is allowed.
		f(w, r)
	}
}
