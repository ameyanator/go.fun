package main

import "fmt"

type EcoStrategy struct{}

func (e *EcoStrategy) drive() {
	fmt.Println("driving in economical mode")
}
