package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc       = bufio.NewScanner(os.Stdin)
	wr       = bufio.NewWriter(os.Stdout)
	L        []int
	J        []int
	knapsack [][]int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	L = make([]int, n)
	J = make([]int, n)
	knapsack = make([][]int, n+1)

	for i := 0; i <= n; i++ {
		knapsack[i] = make([]int, 100)
	}
	for i := 0; i < n; i++ {
		L[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		J[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 100; j++ {
			if j < L[i] {
				knapsack[i+1][j] = knapsack[i][j]
			} else {
				knapsack[i+1][j] = max(knapsack[i][j], knapsack[i][j-L[i]]+J[i])
			}
		}
	}
	fmt.Fprintln(wr, knapsack[n][99])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
