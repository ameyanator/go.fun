package main

type Visitor interface {
	acceptForSquare(*Square)
	acceptForRectangle(*Rectangle)
	acceptForTriangle(*RightTriangle)
}
