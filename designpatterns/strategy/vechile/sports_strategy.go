package main

import "fmt"

type SportsStrategy struct{}

func (s *SportsStrategy) drive() {
	fmt.Println("driving in sports mode")
}
