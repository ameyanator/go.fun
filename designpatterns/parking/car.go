package main

type Car struct {
	licensePlate string
}

func NewCar(licensePlate string) *Car {
	return &Car{
		licensePlate: licensePlate,
	}
}

func (c *Car) getLicensePlateNumber() string {
	return c.licensePlate
}

func (c *Car) getVehicleType() VehicleType {
	return CarType
}
