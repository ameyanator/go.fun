package main

type Truck struct {
	licensePlate string
}

func NewTruck(licensePlate string) *Truck {
	return &Truck{
		licensePlate: licensePlate,
	}
}

func (c *Truck) getLicensePlateNumber() string {
	return c.licensePlate
}

func (c *Truck) getVehicleType() VehicleType {
	return TruckType
}
