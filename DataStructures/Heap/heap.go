package main

import (
	"container/heap"
	"fmt"
)

// Item represents an item in the heap.
type Item[T any] struct {
	Value    T
	Priority int
}

// PriorityQueue implements a min-heap based on the Priority field of Item.
type PriorityQueue[T any] []*Item[T]

// Len returns the number of elements in the heap.
func (pq PriorityQueue[T]) Len() int { return len(pq) }

// Less returns whether the element with index i should sort before the element with index j.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the elements with indexes i and j.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an element to the heap.
func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(*Item[T])
	*pq = append(*pq, item)
}

// Pop removes and returns the minimum element from the heap.
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue[string], 0)

	// Push items onto the heap.
	heap.Push(&pq, &Item[string]{Value: "foo", Priority: 3})
	heap.Push(&pq, &Item[string]{Value: "bar", Priority: 1})
	heap.Push(&pq, &Item[string]{Value: "baz", Priority: 2})

	// Pop items from the heap in order of priority.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%s:%d ", item.Value, item.Priority)
	}
}
