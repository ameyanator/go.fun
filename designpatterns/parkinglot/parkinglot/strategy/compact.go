package strategy

import (
	"fmt"

	"goinpractice.com/designpatterns/parkinglot/parkinglot"
)

// CompactSlotStrategy parks a vehicle in the compact slot (first available compact slot)
type CompactSlotStrategy struct{}

// ParkVehicle implements the CompactSlotStrategy
func (s *CompactSlotStrategy) ParkVehicle(floors []parkinglot.Floor) (int, error) {
	for i := len(floors) - 1; i >= 0; i-- {
		for j := len(floors[i].Slots) - 1; j >= 0; j-- {
			slot := floors[i].Slots[j]
			if !slot.Occupied {
				floors[i].Slots[j].Occupied = true
				return (i * 100) + j + 1, nil // Slot number is calculated based on floor and slot index
			}
		}
	}
	fmt.Println("defaulting strategy")
	// If no compact slots are available, use the first available slot strategy
	return (&FirstAvailableSlotStrategy{}).ParkVehicle(floors)
}
