package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	windowSize              time.Duration
	maxRequestInAWindow     int
	requestsInCurrentWindow int
	lastRequestTime         time.Time
	mu                      sync.Mutex
}

func NewRateLimiter(windowSize time.Duration, maxRequestInAWindow int) *RateLimiter {
	return &RateLimiter{
		windowSize:              windowSize,
		maxRequestInAWindow:     maxRequestInAWindow,
		requestsInCurrentWindow: 0,
		lastRequestTime:         time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Since(rl.lastRequestTime) >= rl.windowSize {
		rl.lastRequestTime = time.Now()
		rl.requestsInCurrentWindow = 0
	}
	rl.requestsInCurrentWindow++

	return rl.requestsInCurrentWindow <= rl.maxRequestInAWindow
}

func main() {
	rateLimiter := NewRateLimiter(time.Second, 5)

	for i := 1; i <= 20; i++ {
		if rateLimiter.Allow() {
			fmt.Printf("Request %d allowed\n", i)
		} else {
			fmt.Printf("Request %d throttled\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
