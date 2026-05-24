package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	a := 42

	b := &a

	fmt.Println(b)  // the address of a that b is pointing to
	fmt.Println(&a) // the address of a itself
	fmt.Println(*b) // the value of a that b is pointing to

	*b = 21

	fmt.Printf("a: %d, *b: %d\n", a, *b)

	x, y := 10, 20

	fmt.Printf("Before swap: x: %d, y: %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("After swap: x: %d, y: %d\n", x, y)
}
