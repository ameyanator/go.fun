package uberlld

import (
	"math"
)

type Location struct {
	x float64
	y float64
}

func (l *Location) GetDistanceFrom(point *Location) float64 {
	// fmt.Println("Getting distance between", l, "and ", point)
	return math.Sqrt((l.x-point.x)*(l.x-point.x) + (l.y-point.y)*(l.y-point.y))
}
