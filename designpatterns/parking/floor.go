package main

import "sync"

type Floor struct {
	slots []*Slot
	mu    sync.Mutex
}

func NewFloor(slots []*Slot) *Floor {
	return &Floor{
		slots: slots,
	}
}

func (f *Floor) emptyFloor() {
	for _, slot := range f.slots {
		slot.emptySlot()
	}
}
