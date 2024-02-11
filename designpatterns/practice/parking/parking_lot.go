package main

import (
	"fmt"
	"time"
)

type ParkVehicleStrategy interface {
	parkVehicle(Vehicle, []*Floor) *Slot
}

type TicketManager struct {
	getTicket map[Vehicle]*Ticket
}

type Ticket struct {
	slot      *Slot
	startTime time.Time
}

func NewTicket(slot *Slot) *Ticket {
	return &Ticket{
		slot:      slot,
		startTime: time.Now(),
	}
}

type ParkingEntrance struct {
	strategy      ParkVehicleStrategy
	floors        []*Floor
	entry         chan Vehicle
	ticketManager TicketManager
	closeChannel  chan struct{}
}

func (e *ParkingEntrance) worker() {
	for {
		select {
		case vehicle := <-e.entry:
			slot := e.strategy.parkVehicle(vehicle, e.floors)
			if slot == nil {
				fmt.Println("Can't park the vehicle")
			}

			e.ticketManager.getTicket[vehicle] = NewTicket(slot)
		case <-e.closeChannel:
			return
		}
	}
}

type ParkingExit struct {
	floors        []*Floor
	exit          chan Vehicle
	ticketManager TicketManager
	closeChannel  chan struct{}
}

func (e *ParkingExit) worker() {
	for {
		select {
		case vehicle := <-e.exit:
			ticket := e.ticketManager.getTicket[vehicle]
			slot := ticket.slot
			slot.emptySlot()
			e.ticketManager.getTicket[vehicle] = nil
		case <-e.closeChannel:
			return
		}
	}
}

type ParkingLot struct {
	floors  []*Floor
	entries []*ParkingEntrance
	exits   []*ParkingExit
}

func NewParkingLot(floors []*Floor, numberOfEntries, numberOfExits int) *ParkingLot {
	entires, exits := make([]*ParkingEntrance, numberOfEntries), make([]*ParkingExit, numberOfExits)

	for _, entry := range entires {
		go entry.worker()
	}

	return &ParkingLot{
		floors:  floors,
		entries: entires,
		exits:   exits,
	}
}
