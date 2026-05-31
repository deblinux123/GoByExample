package main

import "fmt"

func myPanic() {
	panic("There is an Promblem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd, Error:\n", r)
		}
	}()

	myPanic()
	fmt.Println("After myPanic Function")
}
