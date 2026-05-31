package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contain:", s.Contains("Bbak", "a"))
	p("Count", s.Count("my name is farid kaki", "m"))
	p("HasPrefix", s.HasPrefix("test", "st"))
	p("HasPrefix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "st"))
	p("HasSuffix", s.HasSuffix("test", "te"))
	p("Index", s.Index("Farid", "i"))
	p("Upper", s.ToUpper("farid"))
	p("Lower", s.ToLower("FARID"))
	p("Split", s.Split("My-name-is-farid", "-"))

}
