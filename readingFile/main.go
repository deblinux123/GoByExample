package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := filepath.Join(os.TempDir(), "dat")

	dat, err := os.ReadFile(path)

	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(path)
	check(err)
	b := make([]byte, 5)
	n, err := f.Read(b)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(b[:n]))

	o, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)

	n2, err := f.Read(b2)
	check(err)

	fmt.Printf("%d byte: @ %d\n", n2, o)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	o2, err := f.Seek(6, io.SeekStart)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)

	check(err)
	fmt.Printf("%d byte @ %d: %s\n", n3, o2, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)

	check(err)

	fmt.Printf("5 bytes: %s\n", string(b4))

}
