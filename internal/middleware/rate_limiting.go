package middleware

import (
	"fmt"
	"net/http"
)

type RateLimiter interface {
	IsAllowed(string) bool
}

func RateLimiting(rl RateLimiter, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.IsAllowed(r.RemoteAddr) {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, http.StatusText(http.StatusTooManyRequests))
			return
		}
		f(w, r)
	}
}
