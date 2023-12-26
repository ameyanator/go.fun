package main

import "fmt"

type Customer struct {
	id   string
	item Subject
}

func NewCustomer(observable Subject, id string) *Customer {
	fmt.Printf("Customer %s is looking at item %s\n", id, observable.Name())
	return &Customer{
		id:   id,
		item: observable,
	}
}

func (c *Customer) update() {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, c.item.Name())
}

func (c *Customer) getId() string {
	return c.id
}
