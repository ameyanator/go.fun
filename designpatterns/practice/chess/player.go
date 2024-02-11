package main

type Player struct {
	color string
}

func NewPlayer(color string) *Player {
	return &Player{
		color: color,
	}
}
