package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Print("Write ", i, " as 'one'")
	case 2:
		fmt.Print("Write ", i, " as 'two'")
	case 3:
		fmt.Print("Write ", i, " as 'three'")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch b := i.(type) {
		case bool:
			fmt.Printf("I'm a bool\n")
		case int:
			fmt.Printf("I'm an int\n")
		default:
			fmt.Printf("Don't know type %T\n", b)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
