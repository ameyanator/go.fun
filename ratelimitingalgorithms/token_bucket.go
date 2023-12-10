package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens       int           // this is the number of tokens in the bucket at any interval
	capacity     int           // this is the max size of the bucket
	refillRate   time.Duration // rate at which we will refill tokens in the bucket
	stopRefiller chan struct{} // for us to signal to stop refilling tokens
	mu           sync.Mutex    // to handle race conditions
}

func NewTokenBucket(capacity, tokensPerInterval int, refillRate time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:     capacity,
		refillRate:   refillRate,
		stopRefiller: make(chan struct{}),
	}
	// we will refill tokensPerInterval every refillRate time
	go tb.refillTokens(tokensPerInterval)
	return tb
}

func (tb *TokenBucket) refillTokens(tokensPerInterval int) {
	// ticker is a great way to do something repeatedly to know more
	// check this out - https://gobyexample.com/tickers
	ticker := time.NewTicker(tb.refillRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// handle race conditions
			tb.mu.Lock()
			if tb.tokens+tokensPerInterval <= tb.capacity {
				// if we won't exceed the capacity add tokensPerInterval
				// tokens into our bucket
				tb.tokens += tokensPerInterval
			} else {
				// as we cant add more than capacity tokens, set
				// current tokens to bucket's capacity
				tb.tokens = tb.capacity
			}
			tb.mu.Unlock()
		case <-tb.stopRefiller:
			// let's stop refilling
			return
		}
	}
}

func (tb *TokenBucket) TakeTokens() bool {
	// handle race conditions
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// if there are tokens available in the bucket, we take one out
	// in this case request goes through, thus we return true.
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	// in the case where tokens are unavailable, this request won't
	// go through, so we return false
	return false
}

func (tb *TokenBucket) StopRefiller() {
	// close the channel
	close(tb.stopRefiller)
}

func main() {
	tb := NewTokenBucket(10, 5, time.Second)
	// let's give sometime for refillTokens goroutine to spin up
	time.Sleep(900 * time.Millisecond)

	for i := 1; i <= 15; i++ {
		if tb.TakeTokens() {
			// if take tokens returns true, we will let the request go through
			fmt.Printf("Token taken. Remaining tokens: %d\n", tb.tokens)
		} else {
			// if take tokens returns false, we will throttle the request
			fmt.Printf("Not enought tokens. Remaining tokens: %d\n", tb.tokens)
		}
		// Let's wait sometime before sending in the next request
		time.Sleep(150 * time.Millisecond)
	}
	tb.StopRefiller()
}
