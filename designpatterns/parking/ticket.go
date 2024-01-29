package main

import "time"

type Ticket struct {
	slot      *Slot
	startTime time.Time
	cost      int
}

func NewTicket(slot *Slot) *Ticket {
	return &Ticket{
		slot:      slot,
		startTime: time.Now(),
		cost:      0,
	}
}
