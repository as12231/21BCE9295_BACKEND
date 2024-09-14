package utils

import (
	"sync"
	"time"
)

type RateLimiter struct {
	limit    int
	interval time.Duration
	requests map[string]int
	mutex    sync.Mutex
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: interval,
		requests: make(map[string]int),
	}
}

func (rl *RateLimiter) AllowRequest(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	if count, exists := rl.requests[ip]; exists {
		if now.Sub(time.Unix(0, int64(count))) < rl.interval {
			if count >= rl.limit {
				return false
			}
			rl.requests[ip] = count + 1
		} else {
			rl.requests[ip] = 1
		}
	} else {
		rl.requests[ip] = 1
	}
	return true
}
