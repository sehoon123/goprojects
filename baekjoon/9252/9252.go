package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dp [][]int
)

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	a := nextWord()
	b := nextWord()
	n := len(a)
	m := len(b)
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	var result string
	i, j := n, m
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			result = string(a[i-1]) + result
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	fmt.Fprintln(wr, len(result))
	fmt.Fprintln(wr, result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
