package main

type Car struct {
}

func NewCar() *Car {
	return &Car{}
}

func (v *Car) getType() VehicleType {
	return CarType
}
