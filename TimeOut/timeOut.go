package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Result 1"
	}()

	select {
	case resp := <-c1:
		fmt.Println(resp)
	case <-time.After(1 * time.Second):
		fmt.Println("Time Out 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case resp := <-c2:
		fmt.Println(resp)
	case <-time.After(3 * time.Second):
		fmt.Println("Time Out 2")
	}
}
