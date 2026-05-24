package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println("map:", m)

	fmt.Println("get:", m["b"])

	fmt.Println("len:", len(m))

	delete(m, "a")
	fmt.Println("after delete:", m)

	clear(m)
	fmt.Println("after clear:", m)

	_, prs := m["b"]
	fmt.Println("prs:", prs)

	studentScores := map[string]int{"Alice": 85, "Bob": 90, "Charlie": 78}
	fmt.Println("studentScores:", studentScores)

	for name, score := range studentScores {
		fmt.Println("Name:", name, "Score:", score)
	}
}
