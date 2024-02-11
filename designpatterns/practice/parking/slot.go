package main

type VehicleType int

const (
	CarType   VehicleType = 1
	BikeType  VehicleType = 2
	TruckType VehicleType = 3
)

type Slot struct {
	vehicleType VehicleType
	vehicle     Vehicle
}

func NewSlot(vehicleType VehicleType) *Slot {
	return &Slot{
		vehicleType: vehicleType,
	}
}

func (s *Slot) parkVehicle(vehicle Vehicle) {
	s.vehicle = vehicle
}

func (s *Slot) emptySlot() {
	s.vehicle = nil
}
