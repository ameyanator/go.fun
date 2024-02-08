package main

func main() {
	sq := &Square{side: 5}
	rec := &Rectangle{length: 5, width: 10}
	tri := &RightTriangle{base: 5, height: 10}

	area := &AreaVisitor{}

	sq.accept(area)
	rec.accept(area)
	tri.accept(area)

	perimeter := &PerimeterVisitor{}
	sq.accept(perimeter)
	rec.accept(perimeter)
	tri.accept(perimeter)
}
