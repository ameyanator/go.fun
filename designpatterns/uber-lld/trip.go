package uberlld

type Trip struct {
	cab       *Cab
	passenger *Passenger
	From      *Location
	To        *Location
}

func NewTrip() *Trip {
	return &Trip{
		cab: &Cab{
			id:         123,
			driverName: "Ameya",
		},
		passenger: &Passenger{
			name: "Soumya",
			age: 25,
		},
		From: &Location{
			x: 0,
			y: 0,
		},
		To: &Location{
			x: 100,
			y: 150,
		},
	}
}
