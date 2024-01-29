package main

type TicketManager struct {
	getTicket map[Vehicle]*Ticket
}

func NewTicketManager() *TicketManager {
	return &TicketManager{
		getTicket: make(map[Vehicle]*Ticket),
	}
}
