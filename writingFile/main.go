package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d := []byte("hello\ngo\n")
	path := filepath.Join(os.TempDir(), "dat1")

	err := os.WriteFile(path, d, 0644)
	check(err)

	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	defer f.Close()

	d2 := []byte{114, 111, 109, 101, 10}

	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	n3, err := f.WriteString("Testing\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Buffered\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n4)

	w.Flush()
}

// /writingFile$ cat /tmp/dat1
// hello
// go
// /writingFile$ cat /tmp/dat2
// rome
// Testing
// Buffered
