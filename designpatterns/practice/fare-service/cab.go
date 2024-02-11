package main

type Cab struct {
	id          int
	numberPlate string
	driverName  string
}

func NewCab(id int, numberPlate, driverName string) *Cab {
	return &Cab{
		id:          id,
		numberPlate: numberPlate,
		driverName:  driverName,
	}
}
