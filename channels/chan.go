package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "this is a testing text" }()

	msg := <-message

	fmt.Println(msg)

	number := make(chan int)

	go recivedNumbers(10, number)

	for i := 0; i < 10; i++ {
		result := <-number
		fmt.Println(result + i)
	}

}

func recivedNumbers(numRange int, reciver chan int) {
	for i := range numRange {
		reciver <- i
		fmt.Println("Recived: ", reciver)
	}
}
