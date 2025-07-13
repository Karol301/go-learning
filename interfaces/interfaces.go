package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, heigth float64
}

func (r rect) area() float64 {
	return r.width * r.heigth
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.heigth)
}

func getInfo(s shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.area(), s.perimeter())
}

func main() {
	r := rect{width: 5, heigth: 10}
	getInfo(r)
}
