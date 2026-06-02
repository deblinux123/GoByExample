package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://liara.ir")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(strings.Repeat("=", 20), " Response status:", resp.StatusCode, " ", strings.Repeat("=", 20))

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
