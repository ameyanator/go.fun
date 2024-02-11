package main

type Knight struct {
	color string
}

func (p *Knight) getColor() string {
	return p.color
}

func (p *Knight) getPieceName() string {
	return "knight"
}
