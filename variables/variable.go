package main

import "fmt"

func main() {
	var a = "initial variable"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true

	var e int

	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e)

	txt := "short declaration"

	fmt.Printf("this is '%s'\n", txt)
}
