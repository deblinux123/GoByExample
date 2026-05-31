package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "file.txt")

	f := createFile(path)
	defer closeFile(f)
	writingIntoFile(f)
}

func createFile(filePath string) *os.File {
	fmt.Println("Creating file")

	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}
	return f
}

func writingIntoFile(f *os.File) {
	fmt.Println("Writing to file")
	fmt.Fprintf(f, "Data that i want to write into fil")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")

	err := f.Close()

	if err != nil {
		panic(err)
	}
}
