package main

import "fmt"

type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{width: 4, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())

	rPtr := &r

	fmt.Println("Area: ", rPtr.area())
	fmt.Println("Perimeter: ", rPtr.perim())
}
