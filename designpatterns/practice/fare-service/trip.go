package main

import "time"

type Trip struct {
	cab       *Cab
	passenger *Passenger
	from      *Location
	to        *Location
	time      time.Duration
}

func NewTrip(cab *Cab, passenger *Passenger, from, to *Location) *Trip {
	return &Trip{
		cab: cab,
		passenger: passenger,
		to: to,
		from: from,
	}
}

