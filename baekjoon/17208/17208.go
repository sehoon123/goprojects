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

type pair struct {
	burger, fried int
}

type pairs []pair

func (p pairs) Less(i, j int) bool {
	if p[i].burger == p[j].burger {
		return p[i].fried < p[j].fried
	}
	return p[i].burger < p[j].burger
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Len() int {
	return len(p)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n, m, k := nextInt(), nextInt(), nextInt()
	dp := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	order := make(pairs, n+1)
	for i := 1; i < n+1; i++ {
		order[i].burger, order[i].fried = nextInt(), nextInt()
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			for w := 1; w < k+1; w++ {
				if j >= order[i].burger && w >= order[i].fried {
					dp[i][j][w] = max(dp[i-1][j][w], dp[i-1][j-order[i].burger][w-order[i].fried]+1)
				} else {
					dp[i][j][w] = dp[i-1][j][w]
				}
			}
		}
	}

	fmt.Fprintln(wr, dp[n][m][k])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
