package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// variadic function
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// multy return function
func vals(a, b int) (int, int) {

	add := a + b

	sub := a - b

	return add, sub
}

func main() {
	result := add(3, 4)

	fmt.Println("result:", result)

	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Sum:", total)

	added, subtracted := vals(result, total)

	fmt.Println("Added:", added, "Subtracted:", subtracted)
}
