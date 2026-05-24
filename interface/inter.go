package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	redius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with redius: ", c.redius)
	}
}

func main() {
	r := rect{width: 10, height: 18}
	c := circle{redius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
