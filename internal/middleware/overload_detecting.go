package middleware

import (
	"fmt"
	"net/http"
)

type OverloadDetector interface {
	IsOverloaded() bool
}

func OverloadDetecting(od OverloadDetector, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if od.IsOverloaded() {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, http.StatusText(http.StatusServiceUnavailable))
			return
		}
		f(w, r)
	}
}
