package main

type Passenger struct {
	name, phoneNumber string
}

func NewPassenger(name, phoneNumber string) *Passenger {
	return &Passenger{
		name:        name,
		phoneNumber: phoneNumber,
	}
}
