package strategy

import "goinpractice.com/designpatterns/parkinglot/parkinglot"

// ParkingStrategy is an interface for different parking strategies
type ParkingStrategy interface {
	ParkVehicle(floors []parkinglot.Floor) (int, error)
}
