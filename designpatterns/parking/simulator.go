package main

import "time"

func main() {
	slot1 := NewSlot(CarType, 1)
	slot2 := NewSlot(CarType, 2)
	slot3 := NewSlot(TruckType, 3)

	floor1 := NewFloor([]*Slot{slot1, slot2})
	floor2 := NewFloor([]*Slot{slot3})

	parkingLot := NewParkingLot([]*Floor{floor1, floor2}, 1, 2, &BasicStrategy{})

	go parkingLot.openParkingLot()

	tiago := NewCar("1")
	harrier := NewCar("2")
	truck1 := NewTruck("3")

	parkingLot.entries <- tiago
	parkingLot.entries <- truck1

	parkingLot.exits <- tiago

	parkingLot.entries <- harrier

	time.Sleep(10 * time.Second)
}
