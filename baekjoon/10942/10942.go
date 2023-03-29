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

	for i := 1; i <= n; i++ {
		arr[i] = nextInt()
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}

	for i := 1; i < n; i++ {
		if arr[i] == arr[i+1] {
			dp[i][i+1] = 1
		}
	}

	for i := 2; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			if arr[j] == arr[j+i] && dp[j+1][j+i-1] == 1 {
				dp[j][j+i] = 1
			}
		}
	}

	k := nextInt()

	for i := 0; i < k; i++ {
		s, e := nextInt(), nextInt()
		if dp[s][e] == 1 {
			fmt.Fprintln(wr, 1)
		} else {
			fmt.Fprintln(wr, 0)
		}
	}
}
