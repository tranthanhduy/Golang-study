package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	printArea() float64
}

func main() {

	tr := triangle{
		base:   12.8,
		height: 10.4,
	}

	sq := square{
		sideLength: 5.7,
	}

	printShape(tr)
	printShape(sq)

}

func (tr triangle) printArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (s square) printArea() float64 {
	return s.sideLength * s.sideLength
}

func printShape(s shape) {
	fmt.Println(s.printArea())
}
