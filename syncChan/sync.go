package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker 1 start...")
	time.Sleep(time.Second * 2)
	fmt.Println("Done1.")

	done <- true
}

func worker2(done chan bool) {
	fmt.Println("Worker 2 start...")
	time.Sleep(time.Second)
	fmt.Println("Done2.")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)
	// go worker2(done)
	go add(12, 23, done)

	time.Sleep(time.Second * 3)

	<-done
}

func add(a, b int, done chan bool) {
	fmt.Println("Start adding..")
	time.Sleep(time.Second)
	done <- true
	fmt.Println("Result is :", a+b)
}
