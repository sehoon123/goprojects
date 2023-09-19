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

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 3

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]*2 + dp[i-2]
		dp[i] %= 9901
	}

	fmt.Fprintln(wr, dp[n])
}
