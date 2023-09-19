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

	p, n, h := nextInt(), nextInt(), nextInt()
	dp := make([][]int, p+1)
	wait := make([][]int, p+1)
	for i := 0; i <= p; i++ {
		dp[i] = make([]int, h+1)
		wait[i] = make([]int, 0, n)
	}

	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		if b <= h {
			wait[a] = append(wait[a], b)
		}
	}

	for i := 1; i < p+1; i++ {
		for j := 1; j < len(wait[i])+1; j++ {
			for k := 0; k < j; k++ {
				if dp[i][k]+wait[i][j-1] <= h {
					dp[i][j] = max(dp[i][j], dp[i][k]+wait[i][j-1])
				}
			}
		}
	}

	for i := 1; i < p+1; i++ {
		if len(wait[i]) == 0 {
			fmt.Fprintln(wr, i, 0)
		} else {
			fmt.Fprintln(wr, i, 1000*dp[i][len(wait[i])])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
