package main

import (
	"fmt"
	"time"
)

func main() {
	ticke := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticke.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2600 * time.Millisecond)

	ticke.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
