package main

import (
	"fmt"
	"time"
)

func main() {
	//test 1 simple base fare for everyone
	cab := NewCab(123, "123", "ameya")
	passenger := NewPassenger("soumya", "123456789")

	trip := NewTrip(cab, passenger, &Location{0, 0}, &Location{100, 100})

	fareService := NewBaseFare()

	fmt.Println(fareService.getCost(trip))

	//test 2 distance fare for everyone
	distanceFare := NewDistanceFare(fareService)

	fmt.Println(distanceFare.getCost(trip))

	//test 3 time fare for everyone no extra time
	timeFare1 := NewTimeFare(distanceFare)

	fmt.Println(timeFare1.getCost(trip))

	//test 4 time fare for everyone extra time
	trip.time = time.Minute * 2

	fmt.Println(timeFare1.getCost(trip))

	//test 5 only time and base fare
	timeFare2 := NewTimeFare(fareService)

	fmt.Println(timeFare2.getCost(trip))
}
