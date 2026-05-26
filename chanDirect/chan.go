package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	chan1 := make(chan int, len(numbers))
	chan2 := make(chan int, len(numbers))

	for i := range numbers {
		calclatEvenNumbers(chan1, numbers[i])
		showEvenNumbers(chan1, chan2)
		even := <-chan2
		if even != 0 {
			fmt.Printf("%d Is Even\n", even)
		}
	}

}

func calclatEvenNumbers(rec chan<- int, number int) {
	if number%2 == 0 {
		rec <- number
	} else {
		rec <- 0
	}
}

func showEvenNumbers(rec <-chan int, resp chan<- int) {
	number := <-rec
	resp <- number
}
