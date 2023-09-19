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
	arr := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		arr[i] = nextInt()
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 21)
	}

	dp[1][arr[1]] = 1

	for i := 2; i < n; i++ {
		for j := 0; j < 21; j++ {
			if dp[i-1][j] != 0 {
				if j+arr[i] <= 20 {
					dp[i][j+arr[i]] += dp[i-1][j]
				}
				if j-arr[i] >= 0 {
					dp[i][j-arr[i]] += dp[i-1][j]
				}
			}
		}
	}

	//	for _, v := range dp {
	//		fmt.Fprintln(wr, v)
	//	}
	//
	fmt.Fprintln(wr, dp[n-1][arr[n]])
}
