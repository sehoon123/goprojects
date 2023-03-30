package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	graph := make([][]int, n)
	for i := range graph {
		temp := nextString()
		graph[i] = make([]int, m)
		for j := range graph[i] {
			graph[i][j] = int(temp[j] - '0')
			if graph[i][j] == 1 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if graph[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	max := 0
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	fmt.Fprintln(wr, max*max)

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
