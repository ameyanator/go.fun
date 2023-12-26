package main

// Item represents an item in the heap.
type Item[T any] struct {
	Value T
}

// PriorityQueue implements a heap based on the comparison function cmp.
type PriorityQueue[T any] struct {
	Items []Item[T]
	cmp   func(a, b T) bool
}

// NewPriorityQueue creates a new PriorityQueue with the specified comparison function.
func NewPriorityQueue[T any](cmp func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{cmp: cmp}
}

// Len returns the number of elements in the heap.
func (pq *PriorityQueue[T]) Len() int { return len(pq.Items) }

// Less returns whether the element with index i should sort before the element with index j.
func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.cmp(pq.Items[i].Value, pq.Items[j].Value)
}

// Swap swaps the elements with indexes i and j.
func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

// Push adds an element to the heap.
func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(Item[T])
	pq.Items = append(pq.Items, item)
}

// Pop removes and returns the element from the heap.
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := pq.Items
	n := len(old)
	item := old[n-1]
	pq.Items = old[0 : n-1]
	return item
}
