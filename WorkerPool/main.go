package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "start job", j)
		time.Sleep(time.Second * 2)
		fmt.Println("Worker", id, "finished job", j)

		results <- j * 2
	}
}

func main() {
	const numJubs = 15

	jobs := make(chan int, numJubs)
	result := make(chan int, numJubs)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= numJubs; j++ {
		jobs <- j
	}

	for a := 1; a <= numJubs; a++ {
		<-result
	}
}
