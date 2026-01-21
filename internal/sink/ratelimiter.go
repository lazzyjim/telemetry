package sink

import (
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	maxBytesPerSec int64
	currentBytes   int64
}

func NewRateLimiter(maxBytesPerSec int64) *RateLimiter {
	rl := &RateLimiter{maxBytesPerSec: maxBytesPerSec}

	//reset the counter every second
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			atomic.StoreInt64(&rl.currentBytes, 0)
		}
	}()
	return rl
}

// return false if limit is exceeded
func (rl *RateLimiter) TryConsume(n int64) bool {
	newVal := atomic.AddInt64(&rl.currentBytes, n)
	return newVal <= rl.maxBytesPerSec
}
