package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	train []int
	dp    [][]int
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
	train = make([]int, n+1)
	dp = make([][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		train[i] = train[i-1] + nextInt()
	}

	//fmt.Fprintln(wr, train)

	limit := nextInt()

	//for i := 1; i < 4; i++ {
	//	for k := 1; k <= limit; k++ {
	//		for j := i * k; j <= n; j++ {
	//			dp[i][j] = max(dp[i][j-1], dp[i-1][j-k]+train[j]-train[j-k])
	//		}
	//	}
	//}

	for i := 1; i < 4; i++ {
		for j := i * limit; j <= n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j-limit]+train[j]-train[j-limit])
		}
	}

	fmt.Fprintln(wr, dp[3][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
