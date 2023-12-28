package main

type Player struct {
	name  string
	piece Piece
}

func NewPlayer(name string, piece Piece) *Player {
	return &Player{
		name:  name,
		piece: piece,
	}
}
