package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(w Writer) {
	w("hello")
}
func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})

}
