package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
		},
	}

	var wg sync.WaitGroup

	doIncremet := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Go(func() {
		doIncremet("a", 1000)
	})

	wg.Go(func() {
		doIncremet("c", 100)
	})

	wg.Go(func() {
		doIncremet(
			"b", 1020,
		)
	})

	wg.Wait()

	fmt.Println(c.counters)
}
