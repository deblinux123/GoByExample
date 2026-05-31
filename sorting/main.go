package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}

	slices.Sort(strs)

	fmt.Println(strs)

	ints := []int{4, 3, 5, 1, 2}

	slices.Sort(ints)

	fmt.Println(ints)

	s := slices.IsSorted(ints)
	fmt.Println(s)
}
