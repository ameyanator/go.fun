package main

// type Slot struct {
// 	vehicleType VehicleType
// 	occupied    bool
// 	vehicle     Vehicle
// }

// func NewSlot(vehicleType VehicleType) *Slot {
// 	return &Slot{
// 		vehicleType: vehicleType,
// 		occupied:    false,
// 		vehicle:     nil,
// 	}
// }

// func (s *Slot) parkVehicle(v Vehicle) error {
// 	if v.getVehicleType() != s.vehicleType {
// 		return errors.New("Vehicle Type and Slot Type does not match, please try to park in another slot")
// 	}
// 	if s.occupied {
// 		return errors.New("A vehicle is already parked in this slot, please try to park in another slot")
// 	}
// 	s.vehicle = v
// 	s.occupied = true
// 	return nil
// }
