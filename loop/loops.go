package main

import "fmt"

func main() {
	i := 1

	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	for num := range 10 {
		fmt.Println(num)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		}
	}
}
