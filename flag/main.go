package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 34, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string

	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("Word:", *wordPtr)
	fmt.Println("Numb:", *numPtr)
	fmt.Println("Fork:", *forkPtr)
	fmt.Println("Svar:", svar)
	fmt.Println("Tail:", flag.Args())
}
