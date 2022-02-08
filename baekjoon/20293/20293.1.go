package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	R, C := nextInt(), nextInt()
	N := nextInt()

	gasStation := make([][3]int, N+2)
	for i := 1; i < N+1; i++ {
		gasStation[i][0], gasStation[i][1], gasStation[i][2] = nextInt()-1, nextInt()-1, nextInt()
	}
	gasStation[0][0], gasStation[0][1], gasStation[0][2] = 0, 0, 0
	gasStation[N+1][0], gasStation[N+1][1], gasStation[N+1][2] = R-1, C-1, 0

	sort.Slice(gasStation, func(i, j int) bool {
		return gasStation[i][0]+gasStation[i][1] < gasStation[j][0]+gasStation[j][1]
	})

	left, right := 0, 6020
	for left < right {
		dp := make([]int, N+2)
		for i := 0; i < N+2; i++ {
			dp[i] = -1
		}
		mid := (left + right) / 2
		dp[0] = mid
		for i := 1; i <= N+1; i++ {
			for j := 0; j < i; j++ {
				if dp[j] < abs(gasStation[i][0]-gasStation[j][0])+abs(gasStation[i][1]-gasStation[j][1]) || dp[j] == -1 {
					continue
				}
				dp[i] = max(dp[i], dp[j]-(abs(gasStation[i][0]-gasStation[j][0])+abs(gasStation[i][1]-gasStation[j][1]))+gasStation[i][2])
			}
		}
		if dp[N+1] == -1 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Fprintln(wr, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
