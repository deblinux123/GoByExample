package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func main() {

	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{"Bob", 25})

	fmt.Println(Person{name: "Charlie"})

	p := NewPerson("farid", 23)

	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)

	pPtr := &p

	fmt.Println(pPtr)
	fmt.Println(*pPtr)
	fmt.Println((*pPtr).name)
	fmt.Println((*pPtr).age)

	dog := struct {
		name string
		age  int
	}{
		name: "Buddy",
		age:  5,
	}

	fmt.Println(dog)
}
