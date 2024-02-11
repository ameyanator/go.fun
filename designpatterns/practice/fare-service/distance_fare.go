package main

type DistanceFare struct {
	fareService FareService
}

func NewDistanceFare(fareservice FareService) *DistanceFare {
	return &DistanceFare{
		fareService: fareservice,
	}
}

func (f *DistanceFare) getCost(trip *Trip) float64 {
	return trip.from.GetDistanceFrom(trip.to)*4.5 + f.fareService.getCost(trip)
}
