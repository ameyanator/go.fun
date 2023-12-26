package main

type Subject interface {
	register(Observer)
	deregister(Observer)
	Name() string
	notifyAll()
}
