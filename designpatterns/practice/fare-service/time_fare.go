package main

type TimeFare struct {
	fareService FareService
}

func NewTimeFare(fareservice FareService) *TimeFare {
	return &TimeFare{
		fareService: fareservice,
	}
}

func (f *TimeFare) getCost(trip *Trip) float64 {
	// fmt.Println("Trip time is ", trip.time)
	if trip.time == 0 {
		// fmt.Println("Ignoring trip time")
		return f.fareService.getCost(trip)
	}
	return trip.time.Minutes()*2.3 + f.fareService.getCost(trip)
}
