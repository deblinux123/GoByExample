package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))

	dif := now.Sub(then)

	p(dif)
	p(dif.Hours())
	p(dif.Seconds())

	p(then.Add(dif))
	p(then.Add(-dif))
}
