package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(10))
	fmt.Println()

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64() * 5)

	s := rand.NewPCG(42, 1024)
	r := rand.New(s)
	fmt.Println(r.IntN(100))
	fmt.Println(r.IntN(100))

}
