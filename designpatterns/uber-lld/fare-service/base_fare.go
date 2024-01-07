package fareservice

import (
	"fmt"

	uberlld "goinpractice.com/designpatterns/uber-lld"
)

type BaseFare struct {
}

func NewBaseFare() *BaseFare {
	return &BaseFare{}
}

func (b *BaseFare) GetPrice(trip uberlld.Trip) float64 {
	fmt.Println("Calculating Base Fare")
	return 5
}
