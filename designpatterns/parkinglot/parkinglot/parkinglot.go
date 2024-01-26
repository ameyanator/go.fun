// parkinglot/parkinglot.go
package parkinglot

import "errors"

// VehicleType represents the type of vehicle
type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
)

// Vehicle represents a vehicle with a unique identifier and type
type Vehicle struct {
	LicensePlate string
	VehicleType  VehicleType
}

// Slot represents a parking slot
type Slot struct {
	SlotNumber int
	Occupied   bool
	Vehicle    *Vehicle
}

// Floor represents a floor in the parking lot with a set of slots
type Floor struct {
	FloorNumber int
	Slots       []Slot
}

// ParkingStrategy is an interface for different parking strategies
type ParkingStrategy interface {
	ParkVehicle(floors []Floor) (int, error)
}

// ParkingLot represents the entire parking lot with multiple floors and a parking strategy
type ParkingLot struct {
	Floors   []Floor
	Strategy ParkingStrategy
}

// NewParkingLot initializes a parking lot with specified floors, slots per floor, and parking strategy
func NewParkingLot(numFloors, numSlotsPerFloor int, strategy ParkingStrategy) *ParkingLot {
	parkingLot := ParkingLot{
		Strategy: strategy,
	}

	for i := 0; i < numFloors; i++ {
		floor := Floor{
			FloorNumber: i + 1,
			Slots:       make([]Slot, numSlotsPerFloor),
		}
		parkingLot.Floors = append(parkingLot.Floors, floor)
	}

	return &parkingLot
}

// ParkVehicle parks a vehicle in the parking lot using the specified strategy
func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) (int, error) {
	return pl.Strategy.ParkVehicle(pl.Floors)
}

// RemoveVehicle removes a vehicle from the parking lot
func (pl *ParkingLot) RemoveVehicle(slotNumber int) error {
	if slotNumber <= 0 || slotNumber > len(pl.Floors)*len(pl.Floors[0].Slots) {
		return errors.New("invalid slot number")
	}

	// Calculate floor and slot index from the slot number
	floorIndex := (slotNumber - 1) / len(pl.Floors[0].Slots)
	slotIndex := (slotNumber - 1) % len(pl.Floors[0].Slots)

	if !pl.Floors[floorIndex].Slots[slotIndex].Occupied {
		return errors.New("slot is already empty")
	}

	pl.Floors[floorIndex].Slots[slotIndex].Occupied = false
	pl.Floors[floorIndex].Slots[slotIndex].Vehicle = nil

	return nil
}

// SetParkingStrategy sets a new parking strategy for the parking lot
func (pl *ParkingLot) SetParkingStrategy(newStrategy ParkingStrategy) {
	pl.Strategy = newStrategy
}
