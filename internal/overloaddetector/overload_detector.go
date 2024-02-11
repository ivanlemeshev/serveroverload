package overloaddetector

import (
	"context"
	"log"
	"sync/atomic"
	"time"
)

// OverloadDetector implements a simple overload detection algorithm that allows
// to shed load when the system is overloaded.
type OverloadDetector struct {
	checkInterval  time.Duration // interval to check the system load
	overloadFactor time.Duration // factor to determine if the system is overloaded
	isOverloaded   atomic.Bool   // flag to indicate if the system is overloaded
}

// New creates a new OverloadDetector with the given check interval and overload
// factor. The check interval determines how often the system load is checked,
// and the overload factor determines the threshold for the system to be
// considered overloaded.
func New(ctx context.Context, checkInterval, overloadFactor time.Duration) *OverloadDetector {
	od := OverloadDetector{
		checkInterval:  checkInterval,
		overloadFactor: overloadFactor,
	}
	// Start the overload detection algorithm in a separate goroutine.
	go od.run(ctx)
	return &od
}

func (od *OverloadDetector) run(ctx context.Context) {
	ticker := time.NewTicker(od.checkInterval)
	defer ticker.Stop()

	// Start time to measure the elapsed time.
	startTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Check how long it took to process the last batch of requests.
			elapsed := time.Since(startTime)
			log.Println("elapsed: ", elapsed, " overloadFactor: ", od.overloadFactor)
			if elapsed > od.overloadFactor {
				// If it took longer than the overload factor, we're overloaded.
				od.isOverloaded.Store(true)
			} else {
				// Otherwise, we're not overloaded.
				od.isOverloaded.Store(false)
			}
			// Reset the start time.
			startTime = time.Now()
		}
	}
}

func (od *OverloadDetector) IsOverloaded() bool {
	return od.isOverloaded.Load()
}
