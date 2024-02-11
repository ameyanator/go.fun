package main

type Bike struct {
}

func NewBike() *Bike {
	return &Bike{}
}

func (v *Bike) getType() VehicleType {
	return BikeType
}
