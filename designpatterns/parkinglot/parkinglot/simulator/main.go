// main.go
package main

import (
	"fmt"

	"goinpractice.com/designpatterns/parkinglot/parkinglot"
	"goinpractice.com/designpatterns/parkinglot/parkinglot/strategy"
)

func main() {
	// Example usage with FirstAvailableSlotStrategy
	firstAvailableStrategy := &strategy.FirstAvailableSlotStrategy{}
	parkingLot := parkinglot.NewParkingLot(2, 5, firstAvailableStrategy)

	vehicle1 := &parkinglot.Vehicle{LicensePlate: "ABC123", VehicleType: parkinglot.Car}
	vehicle2 := &parkinglot.Vehicle{LicensePlate: "XYZ456", VehicleType: parkinglot.Motorcycle}

	slot1, err1 := parkingLot.ParkVehicle(vehicle1)
	slot2, err2 := parkingLot.ParkVehicle(vehicle2)

	if err1 != nil {
		fmt.Println("Error parking vehicle 1:", err1)
	} else {
		fmt.Println("Vehicle 1 parked at slot:", slot1)
	}

	if err2 != nil {
		fmt.Println("Error parking vehicle 2:", err2)
	} else {
		fmt.Println("Vehicle 2 parked at slot:", slot2)
	}

	// Example usage with CompactSlotStrategy
	compactSlotStrategy := &strategy.CompactSlotStrategy{}
	parkingLot.SetParkingStrategy(compactSlotStrategy)

	vehicle3 := &parkinglot.Vehicle{LicensePlate: "JKL789", VehicleType: parkinglot.Car}
	vehicle4 := &parkinglot.Vehicle{LicensePlate: "MNO012", VehicleType: parkinglot.Motorcycle}

	slot3, err3 := parkingLot.ParkVehicle(vehicle3)
	slot4, err4 := parkingLot.ParkVehicle(vehicle4)

	if err3 != nil {
		fmt.Println("Error parking vehicle 3:", err3)
	} else {
		fmt.Println("Vehicle 3 parked at slot:", slot3)
	}

	if err4 != nil {
		fmt.Println("Error parking vehicle 4:", err4)
	} else {
		fmt.Println("Vehicle 4 parked at slot:", slot4)
	}

	// Display parking lot details
	fmt.Println("\nParking Lot Details:")
	// fmt.Println(parkingLot.GetDetails())
}
