package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "fileName")

	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "fileName"))
	fmt.Println(filepath.Join("dir1/../dir1", "fileName"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir1/file"))
	fmt.Println(filepath.IsAbs("/dir1/file"))

	fileName := "config.json"

	ext := filepath.Ext(fileName)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(fileName, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel)
}
