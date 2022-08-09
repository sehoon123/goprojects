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

	n, m := nextInt(), nextInt()
	coin := make([]int, n+1)

	for i := 1; i <= n; i++ {
		coin[i] = nextInt()
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, m+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coin[i] {
				dp[i][j] = dp[i][j-coin[i]] + dp[i-1][j]
			}
		}
	}

	fmt.Fprintln(wr, dp)
	fmt.Fprintln(wr, dp[n][m])
}
