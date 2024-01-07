package fareservice

import uberlld "goinpractice.com/designpatterns/uber-lld"

type FareService interface {
	GetPrice(trip uberlld.Trip) float64
}
