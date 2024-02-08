package main

import (
	"fmt"
)

type PerimeterVisitor struct{}

func (a *PerimeterVisitor) acceptForSquare(sq *Square) {
	fmt.Println("Perimeter of sq is ", sq.side+sq.side)
}

func (a *PerimeterVisitor) acceptForRectangle(r *Rectangle) {
	fmt.Println("Perimeter of sq is ", 2*(r.length+r.width))
}

func (a *PerimeterVisitor) acceptForTriangle(r *RightTriangle) {
	fmt.Println("Perimeter of sq is ", r.base*2+r.height)
}
