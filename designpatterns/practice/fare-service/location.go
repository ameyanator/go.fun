package main

import "math"

type Location struct {
	lat int
	lon int
}

func NewLocation(lat, lon int) *Location {
	return &Location{
		lat: lat,
		lon: lon,
	}
}

func (l *Location) GetDistanceFrom(to *Location) float64 {
	return math.Abs(float64(l.lat)-float64(to.lat)) + math.Abs(float64(l.lon)-float64(to.lon))
}
