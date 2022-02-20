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
	cards []int
	cache []int
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
	cards = make([]int, n)
	for i := 0; i < n; i++ {
		cards[i] = nextInt()
	}

	fmt.Fprintln(wr, solve(n))

}

func solve(n int) int {
	if n == 1 {
		return cards[0]
	}
	if cache[n] != 0 {
		return cache[n]
	}
	for i := 0; i < n; i++ {
		cache[n] = max(cache[n], solve(i)+cards[n-i-1])
	}
	return cache[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
