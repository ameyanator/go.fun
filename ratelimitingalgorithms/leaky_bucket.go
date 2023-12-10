package main

import (
	"fmt"
	"sync"
	"time"
)

// Representation of incoming requests
type Request struct {
	ID int // ID of our request
	// We should ideally be storing much more values
	// Some examples are URL and Body of the request
}

type LeakyBucket struct {
	queue        []Request     //this represents our bucket - any new request will be sent in this queue
	capacity     int           //this is the capacity of our bucket. We will throttle requests if the bucket is full
	emptyRate    time.Duration //this represents how often we will take requests from our bucket and send to servers
	stopRefiller chan struct{} //we have this variable to signal in case we want to shut down
	mu           sync.Mutex    //this is to handle data race conditions
}

func NewLeakyBucket(capacity int, emptyRate time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		capacity:     capacity,
		emptyRate:    emptyRate,
		stopRefiller: make(chan struct{}),
	}

	//to minic constant dripping we will start a goroutine
	//which will remove requests from the queue at a constant rate
	go lb.removeRequestsFromQueue()
	return lb
}

func (lb *LeakyBucket) removeRequestsFromQueue() {
	// this mimics infinite while loop
	ticker := time.NewTicker(lb.emptyRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// we take a lock to prevent data race conditions
			lb.mu.Lock()
			// if there are requests in the queue
			if len(lb.queue) > 0 {
				fmt.Printf("Serving %d requests\n", len(lb.queue))
				// here we send requests to the server
				// after sending all requests to the server
				// we empty the queue
				lb.queue = nil
			}
			lb.mu.Unlock()
		case <-lb.stopRefiller:
			// when we get this signal we will stop this goroutine
			return
		}
	}
}

func (lb *LeakyBucket) AddRequest(req Request) bool {
	// we take a lock to prevent race conditions
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// if its possible to add into our queue
	// we add the request to the queue
	if len(lb.queue) < lb.capacity {
		lb.queue = append(lb.queue, req)
		return true
	}
	// if its not possible to add request to
	// the queue then we return false
	return false
}

func (lb *LeakyBucket) StopRefiller() {
	close(lb.stopRefiller)
}

func main() {
	// We create a bucket with capacity 3 and requests
	// get sent to the server every 1 second.
	lb := NewLeakyBucket(3, time.Second)

	for i := 1; i <= 5; i++ {
		// make a new request
		req := Request{ID: i}
		if lb.AddRequest(req) {
			// this request was added to the bucket
			fmt.Printf("Request %d added to the queue\n", i)
		} else {
			// this request couldn't be added
			fmt.Printf("Request %d throttled\n", i)
		}
		// some time between requests
		time.Sleep(400 * time.Millisecond)
	}
	lb.StopRefiller()
}
