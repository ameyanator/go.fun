package main

import (
	"fmt"

	uberlld "goinpractice.com/designpatterns/uber-lld"
	fareservice "goinpractice.com/designpatterns/uber-lld/fare-service"
)

func main() {
	trip := uberlld.NewTrip()
	baseFare := fareservice.NewBaseFare()
	distanceFare := fareservice.NewDistanceFare(baseFare)
	timeFare := fareservice.NewTimeFare(distanceFare)

	// fmt.Println("Price with distance ", distanceFare.GetPrice(*trip))
	// fmt.Println()
	fmt.Println("Price with distance and time", timeFare.GetPrice(*trip))
}
