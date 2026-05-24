package main

import "fmt"

func isSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := isSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := isSeq()

	fmt.Println(newInts())
}
