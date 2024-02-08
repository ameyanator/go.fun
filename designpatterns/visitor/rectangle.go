package main

type Rectangle struct {
	length, width int
}

func (s *Rectangle) getType() string {
	return "Rectangle"
}

func (s *Rectangle) accept(v Visitor) {
	v.acceptForRectangle(s)
}
