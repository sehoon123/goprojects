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

	n := nextInt()
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nextInt()
	}

	m := nextInt()
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		fmt.Fprintln(wr, arr[b]-arr[a-1])
	}
}
