package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("ex21.2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	fmt.Fprintln(f, "hello")
}