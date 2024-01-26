package strategy

import (
	"errors"

	"goinpractice.com/designpatterns/parkinglot/parkinglot"
)

// FirstAvailableSlotStrategy parks a vehicle in the first available slot
type FirstAvailableSlotStrategy struct{}

// ParkVehicle implements the FirstAvailableSlotStrategy
func (s *FirstAvailableSlotStrategy) ParkVehicle(floors []parkinglot.Floor) (int, error) {
	for i, floor := range floors {
		for j, slot := range floor.Slots {
			if !slot.Occupied {
				floors[i].Slots[j].Occupied = true
				return (i * 100) + j + 1, nil // Slot number is calculated based on floor and slot index
			}
		}
	}
	return 0, errors.New("no available slots")
}
