package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "this is a string"

	h := sha256.New()

	h.Write([]byte(s))

	fmt.Println(h.BlockSize())
	fmt.Println(h.Size())

	bs := h.Sum(nil)
	fmt.Println(s)

	fmt.Printf("%x\n", bs)

	pass := []byte("12345")

	passHash := sha256.New()
	passHash.Write(pass)
	bsPass := passHash.Sum(nil)
	fmt.Printf("%x\n", bsPass)
}
