package main

import "fmt"

func main() {
	var s []string

	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)

	fmt.Println("init:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("copy:", c)

	l := s[2:4]
	fmt.Println("slice1:", l)

	numbers := make([]int, 10)

	for i := range numbers {
		numbers[i] = i * 2
	}

	fmt.Println(numbers)

	names := []string{"farid", "sami", "mohamed", "ahmed"}

	for i, name := range names {
		fmt.Println("Index:", i, "Name:", name)
	}
}
