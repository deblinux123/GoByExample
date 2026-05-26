package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("Recived:", msg1)
		case msg2 := <-c2:
			fmt.Println("Recived:", msg2)
		}
	}
}
