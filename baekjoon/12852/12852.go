package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	cache  []int
	before []int
)

func main() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Scan(&n)
	cache = make([]int, n+1)

	fmt.Fprintln(wr, solve(n))
	for {
		fmt.Fprintf(wr, "%d ", n)
		if n == 1 {
			break
		}
		if n%2 == 0 && cache[n] == cache[n/2]+1 {
			n /= 2
			continue
		} else if n%3 == 0 && cache[n] == cache[n/3]+1 {
			n /= 3
			continue
		} else {
			n -= 1
			continue
		}
	}
}

func solve(n int) int {
	if n == 1 {
		return 0
	}

	if cache[n] != 0 {
		return cache[n]
	}

	A := solve(n-1) + 1
	B := int(1e9)
	C := int(1e9)
	if n%2 == 0 {
		B = solve(n/2) + 1
	}
	if n%3 == 0 {
		C = solve(n/3) + 1
	}

	cache[n] = min(A, min(B, C))

	return cache[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
