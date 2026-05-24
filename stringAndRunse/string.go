package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("String:", s)

	fmt.Println("len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	txt := "This is a string"

	for idx, runValue := range txt {
		fmt.Printf("%#U starts at byte position %d\n", runValue, idx)
	}

	fmt.Print("\nUsing DocodeRuneInString: ")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width

		examineRuner(runeValue)

	}
}

func examineRuner(r rune) {
	switch r {
	case 't':
		fmt.Println("Found a t")
	case 'a':
		fmt.Println("Found an a")
	}
}
