package main

import "fmt"

type BasicStrategy struct{}

func (b *BasicStrategy) findSlot(floors []*Floor, vehicle Vehicle) *Slot {
	for _, floor := range floors {
		floor.mu.Lock()
		defer floor.mu.Unlock()
		for _, slot := range floor.slots {
			if slot.occupied == false && slot.vehicleType == vehicle.getVehicleType() {
				slot.occupied = true
				slot.vehicle = vehicle
				return slot
			}
		}
	}
	fmt.Println("Couldn't find any slot returning nil")
	return nil
}
