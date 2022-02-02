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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	r, c := nextInt(), nextInt()
	graph := make([]string, r)
	for i := 0; i < r; i++ {
		sc.Scan()
		graph[i] = sc.Text()
	}
	fmt.Fprintln(wr, c, graph)
}
