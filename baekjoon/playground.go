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
	b := sc.Bytes()
	r, f := 0, 1

	for _, c := range b {
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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a, b, c := nextInt(), nextInt(), nextInt()
	fmt.Fprintln(wr, a, b, c)
}
