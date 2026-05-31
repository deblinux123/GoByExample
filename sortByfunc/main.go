package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	names := []string{"farid", "babak", "ahmad", "ali"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(names, lenCmp)

	fmt.Println(names)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Farid", age: 32},
		Person{name: "babak", age: 31},
		Person{name: "Asma", age: 29},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)
}
