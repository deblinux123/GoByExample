package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	for i := range 3 {
		messages <- fmt.Sprintf("Message %d", i)
		fmt.Println(<-messages)
	}

}
