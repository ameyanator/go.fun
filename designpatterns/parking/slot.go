package main

import "errors"

type Slot struct {
	vehicleType VehicleType
	occupied    bool
	vehicle     Vehicle
	id          int
}

func NewSlot(vehicleType VehicleType, id int) *Slot {
	return &Slot{
		vehicleType: vehicleType,
		occupied:    false,
		vehicle:     nil,
		id:          id,
	}
}

func (s *Slot) parkVehicle(v Vehicle) error {
	if v.getVehicleType() != s.vehicleType {
		return errors.New("Vehicle Type and Slot Type does not match, please try to park in another slot")
	}
	if s.occupied {
		return errors.New("A vehicle is already parked in this slot, please try to park in another slot")
	}
	s.vehicle = v
	s.occupied = true
	return nil
}

func (s *Slot) emptySlot() {
	s.occupied = false
	s.vehicle = nil
}
