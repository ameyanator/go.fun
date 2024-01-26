package main

import (
	"sync"
	"time"
)

type SlidingWindowRateLimiter struct {
	queue      []time.Time
	capacity   int
	windowSize time.Duration
	mu         sync.Mutex
}

func NewRateLimiter(capacity int, windowSize time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		capacity:   capacity,
		windowSize: windowSize,
	}
}
