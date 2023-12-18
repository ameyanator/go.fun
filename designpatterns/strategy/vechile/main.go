package main

func main() {
	sportsStrategy := &SportsStrategy{}
	ecoStrategy := &EcoStrategy{}

	vehicle := New(ecoStrategy)
	vehicle.drive()
	vehicle.changeStrategy(sportsStrategy)
	vehicle.drive()
}
