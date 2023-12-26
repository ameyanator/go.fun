package main

type Request struct {
	currentFloor int
	desiredFloor int
	direction    Direction
	location     Location
}

func NewRequest(currentFloor, desiredFloor int, direction Direction, location Location) *Request {
	return &Request{
		currentFloor: currentFloor,
		desiredFloor: desiredFloor,
		direction:    direction,
		location:     location,
	}
}
