package middleware

import (
	"fmt"
	"net/http"
)

// OverloadDetector is an interface that allows to detect if the system is
// overloaded.
type OverloadDetector interface {
	IsOverloaded() bool
}

// OverloadDetecting is a middleware that detects if the system is overloaded.
// If the system is overloaded, it responds with a status code 503. Otherwise,
// it calls the original handler.
func OverloadDetecting(od OverloadDetector, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the server is overloaded.
		if od.IsOverloaded() {
			// Response with status code 503
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, http.StatusText(http.StatusServiceUnavailable))
			return
		}
		// Call the original handler if the request is allowed.
		f(w, r)
	}
}
