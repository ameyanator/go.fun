package main

import (
	"fmt"
	"sync"
)

// https://leetcode.com/problems/design-bounded-blocking-queue/description/

type BoundedBlockingQueue struct {
	elements   []int
	mu         sync.Mutex
	capacity   int
	cond       *sync.Cond
	returnChan chan int
}

func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	bq := &BoundedBlockingQueue{
		elements:   make([]int, capacity),
		capacity:   capacity,
		returnChan: make(chan int, 20),
	}
	bq.cond = sync.NewCond(&bq.mu)
	return bq
}

func (b *BoundedBlockingQueue) enqueue(val int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.elements) >= b.capacity {
		b.cond.Wait()
	}

	b.elements = append(b.elements, val)
	b.cond.Broadcast()
}

func (b *BoundedBlockingQueue) dequeue() {
	b.mu.Lock()
	defer b.mu.Lock()

	for len(b.elements) == 0 {
		b.cond.Wait()
	}
	val := b.elements[0]
	b.elements = b.elements[1:]
	b.returnChan <- val
}

func (b *BoundedBlockingQueue) size() int {
	return len(b.elements)
}

func main() {
	queue := NewBoundedBlockingQueue(3)

	for i := 0; i < 5; i++ {
		go queue.enqueue(i)
	}

	for i := 0; i < 4; i++ {
		go queue.dequeue()
	}

	// fmt.Println("size of queue is ", queue.size())

	for i := 5; i < 10; i++ {
		go queue.enqueue(i)
	}

	for i := 4; i < 10; i++ {
		 queue.dequeue()
	}
	counter := 0
	for {
		select {
		case val := <-queue.returnChan:
			fmt.Println("Received from queue ", val)
			counter++
			if counter == 10 {
				break;
			}
		}
	}
}