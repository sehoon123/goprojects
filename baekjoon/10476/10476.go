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

	graph := make([][]int, 201)
	dp := make([][][]int, 202)
	n, k := nextInt(), nextInt()

	for i := 1; i <= n+1; i++ {
		graph[i] = make([]int, 2)
		a, b := nextInt(), nextInt()
		graph[i][0] = a
		graph[i][1] = b

		dp[i] = make([][]int, 202)
		for j := 0; j < 201; j++ {
			dp[i][j] = make([]int, 3)
		}
	}

	//dp[floor][closed][now_state]
	// now_state: 0: none, 1: left, 2: right
	dp[1][0][0] = graph[1][0] + graph[1][1]
	dp[1][1][1] = graph[1][1]
	dp[1][1][2] = graph[1][0]

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j > i {
				break
			}
			if j >= 1 {
				dp[i][j][1] = max(dp[i-1][j-1][0], dp[i-1][j-1][1]) + graph[i][1]
				dp[i][j][2] = max(dp[i-1][j-1][0], dp[i-1][j-1][2]) + graph[i][0]
			}
			if i != j {
				dp[i][j][0] = threeMax(dp[i-1][j][0], dp[i-1][j][1], dp[i-1][j][2]) + graph[i][0] + graph[i][1]
			}
		}
	}
	fmt.Fprintln(wr, threeMax(dp[n][k][0], dp[n][k][1], dp[n][k][2]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func threeMax(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
