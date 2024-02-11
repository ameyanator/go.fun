package main

type Floor struct {
	slots []*Slot
}

func NewFloor(slots []*Slot) *Floor {
	return &Floor{
		slots: slots,
	}
}
