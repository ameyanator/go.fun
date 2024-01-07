package fareservice

import (
	"fmt"

	uberlld "goinpractice.com/designpatterns/uber-lld"
)

type DistanceFare struct {
	fareService FareService
}

func NewDistanceFare(fareService FareService) *DistanceFare {
	return &DistanceFare{fareService: fareService}
}

func (d *DistanceFare) GetPrice(trip uberlld.Trip) float64 {
	fmt.Println("Calculating Distance Fare for distance ", trip.From.GetDistanceFrom(trip.To))
	return 10*trip.From.GetDistanceFrom(trip.To) + d.fareService.GetPrice(trip)
}
