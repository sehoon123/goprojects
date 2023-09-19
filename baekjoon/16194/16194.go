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

	n := nextInt()
	cards := make([]int, n+1)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cards[i] = nextInt()
		dp[i] = cards[i]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j]+cards[j])
		}
	}

	fmt.Fprintln(wr, dp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
