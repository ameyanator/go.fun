package main

import "fmt"

func main() {
	elevator := NewElevator(0)

	elevator.sendUpRequest(NewRequest(5, -1, UP, OUTSIDE_ELEVATOR))
	elevator.run()
	fmt.Println("Elevator is at floor ", elevator.currentFloor)
	elevator.sendUpRequest(NewRequest(4, -1, UP, OUTSIDE_ELEVATOR))
	elevator.sendDownRequest(NewRequest(4, 2, DOWN, INSIDE_ELEVATOR))
	elevator.sendUpRequest(NewRequest(4, 5, UP, INSIDE_ELEVATOR))
	elevator.sendUpRequest(NewRequest(3, -1, DOWN, OUTSIDE_ELEVATOR))
	elevator.run()
}
