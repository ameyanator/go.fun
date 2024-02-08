package main

type RightTriangle struct {
	base, height int
}

func (s *RightTriangle) getType() string {
	return "RightTriangle"
}

func (s *RightTriangle) accept(v Visitor) {
	v.acceptForTriangle(s)
}
