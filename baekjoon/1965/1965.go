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
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		arr[i] = nextInt()
		dp[i] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dp[i])
	}

	fmt.Fprintln(wr, ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
