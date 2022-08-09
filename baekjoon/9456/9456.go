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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	T := nextInt()

	for i := 0; i < T; i++ {
		n := nextInt()
		graph := make([][]int, 2)
		dp := make([][]int, 2)
		for j := 0; j < 2; j++ {
			graph[j] = make([]int, n)
			dp[j] = make([]int, n+1)
			for k := 0; k < n; k++ {
				graph[j][k] = nextInt()
			}
		}

		dp[0][1] = graph[0][0]
		dp[1][1] = graph[1][0]

		for j := 2; j <= n; j++ {
			dp[0][j] = max(dp[1][j-1]+graph[0][j-1], max(dp[0][j-2]+graph[0][j-1], dp[1][j-2]+graph[0][j-1]))
			dp[1][j] = max(dp[0][j-1]+graph[1][j-1], max(dp[1][j-2]+graph[1][j-1], dp[0][j-2]+graph[1][j-1]))
		}

		fmt.Fprintln(wr, max(dp[0][n], dp[1][n]))

	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
