package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	n := nextInt()
	fmt.Println(n)
}
