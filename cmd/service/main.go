package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ivanlemeshev/serveroverload/internal/middleware"
	"github.com/ivanlemeshev/serveroverload/internal/password"
	"github.com/ivanlemeshev/serveroverload/internal/ratelimiter"
)

func main() {
	rl := ratelimiter.NewTokenBucket(1, 1)
	http.HandleFunc("GET /password/{length}", middleware.RateLimiting(rl, handler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.PathValue("length"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}
	pwd := password.Generate(length)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, pwd)
}
