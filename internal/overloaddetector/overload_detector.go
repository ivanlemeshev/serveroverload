package overloaddetector

import (
	"context"
	"sync/atomic"
	"time"
)

type OverloadDetector struct {
	checkInterval  time.Duration
	overloadFactor time.Duration
	isOverloaded   atomic.Bool
}

func New(ctx context.Context, checkInterval, overloadFactor time.Duration) *OverloadDetector {
	od := OverloadDetector{
		checkInterval:  checkInterval,
		overloadFactor: overloadFactor,
	}
	go od.run(ctx)
	return &od
}

func (od *OverloadDetector) IsOverloaded() bool {
	return od.isOverloaded.Load()
}

func (od *OverloadDetector) run(ctx context.Context) {
	ticker := time.NewTicker(od.checkInterval)
	defer ticker.Stop()

	checkTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if time.Since(checkTime) > od.overloadFactor {
				od.isOverloaded.Store(true)
			} else {
				od.isOverloaded.Store(false)
			}
			checkTime = time.Now()
		}
	}
}
