package main

import (
	"container/heap"
	"fmt"
)

type Elevator struct {
	currentFloor int
	direction    Direction
	upQueue      *PriorityQueue[*Request]
	downQueue    *PriorityQueue[*Request]
}

func NewElevator(currentFloor int) *Elevator {
	return &Elevator{
		currentFloor: currentFloor,
		direction:    IDLE,
		upQueue: NewPriorityQueue(func(a, b *Request) bool {
			return a.desiredFloor < b.desiredFloor
		}),
		downQueue: NewPriorityQueue(func(a, b *Request) bool {
			return a.desiredFloor > b.desiredFloor
		}),
	}
}

func (e *Elevator) sendUpRequest(upRequest *Request) {
	if upRequest.location == OUTSIDE_ELEVATOR {
		heap.Push(e.upQueue, Item[*Request]{Value: NewRequest(upRequest.currentFloor, upRequest.currentFloor, UP, OUTSIDE_ELEVATOR)})
		fmt.Println("Append up request going to floor ", upRequest.currentFloor)
		return
	}
	heap.Push(e.upQueue, Item[*Request]{Value: upRequest})
	fmt.Println("Append up request going to floor ", upRequest.desiredFloor)
}

func (e *Elevator) sendDownRequest(downRequest *Request) {
	if downRequest.location == OUTSIDE_ELEVATOR {
		heap.Push(e.downQueue, Item[*Request]{Value: NewRequest(downRequest.currentFloor, downRequest.currentFloor, DOWN, OUTSIDE_ELEVATOR)})
		fmt.Println("Append down request going to floor ", downRequest.currentFloor)
		return
	}
	heap.Push(e.downQueue, Item[*Request]{Value: downRequest})
	fmt.Println("Append down request going to floor ", downRequest.desiredFloor)
}

func (e *Elevator) run() {
	for len(e.upQueue.Items) > 0 || len(e.downQueue.Items) > 0 {
		e.processRequests()
	}
	fmt.Println("Elevator is IDLE now")
	e.direction = IDLE
}

func (e *Elevator) processRequests() {
	if e.direction == UP || e.direction == IDLE {
		e.processUpRequests()
		e.processDownRequests()
	} else {
		e.processDownRequests()
		e.processUpRequests()
	}
}

func (e *Elevator) processUpRequests() {
	for len(e.upQueue.Items) > 0 {
		upRequest := e.upQueue.Pop().(Item[*Request])
		e.currentFloor = upRequest.Value.desiredFloor
		fmt.Println("Processing up requests, Elevator is at floor ", e.currentFloor)
	}
	if len(e.downQueue.Items) > 0 {
		e.direction = DOWN
	} else {
		e.direction = IDLE
	}
}

func (e *Elevator) processDownRequests() {
	for len(e.downQueue.Items) > 0 {
		downRequest := e.downQueue.Pop().(Item[*Request])
		e.currentFloor = downRequest.Value.desiredFloor
		fmt.Println("Processing down requests, Elevator is at floor", e.currentFloor)
	}
	if len(e.upQueue.Items) > 0 {
		e.direction = UP
	} else {
		e.direction = IDLE
	}
}
