package main

import (
	"bufio"
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

	dp := make([]int, 100001)
	for i := 1; i <= 100000; i++ {
		dp[i] = 100001
	}

	dp[2] = 1
	dp[4] = 2
	dp[5] = 1

	for i := 6; i <= n; i++ {
		dp[i] = min(dp[i-2], dp[i-5]) + 1
	}

	if dp[n] == 100001 {
		wr.WriteString("-1")
	} else {
		wr.WriteString(strconv.Itoa(dp[n]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
