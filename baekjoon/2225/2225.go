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

	n, k := nextInt(), nextInt()
	dp := make([]int, n+1)

	for i := 1; i <= k; i++ {
		if i == 1 {
			for j := 0; j <= n; j++ {
				dp[j] = 1
			}
			continue
		}
		for j := 0; j <= n; j++ {
			if j == 0 {
				dp[j] = 1
				continue
			}
			dp[j] = (dp[j] + dp[j-1]) % 1000000000
		}
	}

	fmt.Fprintln(wr, dp[n])
}
