package main

import "fmt"

func main() {
	var a [5]int

	fmt.Printf("emp: %v\n", a)

	a[0] = 2
	a[1] = 3
	a[2] = 4
	a[3] = 5
	a[4] = 6

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [3]string{"a", "b", "c"}

	fmt.Println("dcl:", b)

	b = [...]string{"d", "e", "f"}

	fmt.Println("dcl:", b)

	var twoD [2][3]int

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("2d: ", twoD)
}
