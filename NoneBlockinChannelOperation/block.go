package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Recived message:", msg)
	default:
		fmt.Println("No messagee recived.")
	}

	msg := "Hello this is a secret message"

	select {
	case messages <- msg:
		fmt.Println("Senet message:", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Recived message", msg)
	case sig := <-signals:
		fmt.Println("Recived signals:", sig)
	default:
		fmt.Println("No activity")
	}
}
