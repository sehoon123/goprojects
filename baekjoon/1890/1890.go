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
	//board  [][]int
	dx, dy []int
	count  int
	n      int
	//dp     [][]int
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
	board := make([][]int, n)
	dp := make([][]int, n)
	count = 0

	dx = []int{0, 1}
	dy = []int{1, 0}

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			board[i][j] = nextInt()
		}
	}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 || (i == n-1 && j == n-1) {
				continue
			}

			dist := board[i][j]
			down := i + dist
			right := j + dist

			if down < n {
				dp[down][j] += dp[i][j]
			}
			if right < n {
				dp[i][right] += dp[i][j]
			}
		}
	}

	fmt.Fprintln(wr, dp[n-1][n-1])
}
