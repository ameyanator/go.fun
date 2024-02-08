package main

import "fmt"

type AreaVisitor struct{}

func (a *AreaVisitor) acceptForSquare(sq *Square) {
	fmt.Println("Area of sq is ", sq.side*sq.side)
}

func (a *AreaVisitor) acceptForRectangle(r *Rectangle) {
	fmt.Println("Area of sq is ", r.length*r.width)
}

func (a *AreaVisitor) acceptForTriangle(r *RightTriangle) {
	fmt.Println("Area of sq is ", float64(r.base*r.height)*0.5)
}

