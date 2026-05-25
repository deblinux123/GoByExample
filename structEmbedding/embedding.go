package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 12,
		},
		str: "Hello",
	}

	fmt.Printf("Container: {num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("Also num: ", co.base.num)

	fmt.Println("Describe: ", co.describe())

	type describe interface {
		describe() string
	}

	var d describe = co

	fmt.Println("Describe:", d.describe())
}
