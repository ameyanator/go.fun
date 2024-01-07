package fareservice

import (
	"fmt"

	uberlld "goinpractice.com/designpatterns/uber-lld"
)

type TimeFare struct {
	fareService FareService
}

func NewTimeFare(fareService FareService) *TimeFare {
	return &TimeFare{fareService: fareService}
}

func (t *TimeFare) GetPrice(trip uberlld.Trip) float64 {
	fmt.Println("Calculating Time Fare")
	return 5*(trip.From.GetDistanceFrom(trip.To)*3) + t.fareService.GetPrice(trip)
}
