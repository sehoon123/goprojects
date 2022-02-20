package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	n     int
	cache []int
	wine  []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n = nextInt()
	cache = make([]int, n+1)
	wine = make([]int, n)
	for i := 0; i < n; i++ {
		wine[i] = nextInt()
	}

	fmt.Fprintln(wr, solve(n))
}

func solve(n int) int {
	if n == 1 {
		return wine[0]
	} else if n == 2 {
		return wine[0] + wine[1]
	} else if n == 3 {
		return max(max(wine[0]+wine[2], wine[1]+wine[2]), wine[0]+wine[1])
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = max(max(solve(n-2), solve(n-3)+wine[n-2])+wine[n-1], solve(n-1))
	return cache[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
