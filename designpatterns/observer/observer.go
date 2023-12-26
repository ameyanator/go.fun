package main

type Observer interface {
	update()
	getId() string
}
