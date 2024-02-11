package main

type BaseFare struct {
}

func NewBaseFare() *BaseFare {
	return &BaseFare{}
}

func (f *BaseFare) getCost(trip *Trip) float64 {
	return 10.0
}
