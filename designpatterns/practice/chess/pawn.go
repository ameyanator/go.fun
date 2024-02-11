package main

type Pawn struct {
	color string
}

func (p *Pawn) getColor() string {
	return p.color
}

func (p *Pawn) getPieceName() string {
	return "pawn"
}
