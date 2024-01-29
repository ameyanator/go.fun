package main

import "fmt"

type ParkingLotEntry struct {
	floors   []*Floor
	strategy FindSlotStrategy
}

type ParkingLotExit struct {
	parkingLot *ParkingLot
}

type ParkingLot struct {
	floors          []*Floor
	entries         chan Vehicle
	exits           chan Vehicle
	closeParkingLot chan struct{}
	strategy        FindSlotStrategy
	ticketManager   *TicketManager
}

func NewParkingLot(floors []*Floor, totalEntries, totalExits int, strategy FindSlotStrategy) *ParkingLot {
	return &ParkingLot{
		floors:          floors,
		entries:         make(chan Vehicle, totalEntries),
		closeParkingLot: make(chan struct{}),
		strategy:        strategy,
		ticketManager:   NewTicketManager(),
		exits:           make(chan Vehicle, totalExits),
	}
}

func (p *ParkingLot) openParkingLot() {
	for {
		select {
		case vehicle, ok := <-p.entries:
			if !ok {
				fmt.Println("Parking lot entries are not working pls have a look, closing lot")
				return
			}
			slot := p.strategy.findSlot(p.floors, vehicle)
			fmt.Println("Vehicle Parked at ", slot)
			ticket := NewTicket(slot)
			p.ticketManager.getTicket[vehicle] = ticket

		case vehicle, _ := <-p.exits:
			ticket := p.ticketManager.getTicket[vehicle]
			slot := ticket.slot
			slot.emptySlot()
			fmt.Println("Vehicle ", vehicle, " has emptied the slot it was occupying")

		case <-p.closeParkingLot:
			fmt.Println("we are closing the parking lot now")
			for _, floor := range p.floors {
				floor.emptyFloor()
			}
			return
		}
	}
}
