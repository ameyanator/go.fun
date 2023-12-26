package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update()
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLen := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getId() == observer.getId() {
			observerList[observerListLen-1], observerList[i] = observerList[i], observerList[observerListLen-1]
			return observerList[:observerListLen-1]
		}
	}
	return observerList
}

func (i *Item) Name() string {
	return i.name
}
