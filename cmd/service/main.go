package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ivanlemeshev/serveroverload/internal/middleware"
	"github.com/ivanlemeshev/serveroverload/internal/overloaddetector"
	"github.com/ivanlemeshev/serveroverload/internal/password"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	od := overloaddetector.New(ctx, 10*time.Microsecond, 11*time.Millisecond)
	http.HandleFunc("GET /password/{length}", middleware.OverloadDetecting(od, handler))
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
