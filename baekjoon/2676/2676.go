package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	T := nextInt()
	for i := 0; i < T; i++ {
		a, b := nextInt(), nextInt()
		fmt.Fprintln(wr, 1+(a-b)*b)
	}
}
