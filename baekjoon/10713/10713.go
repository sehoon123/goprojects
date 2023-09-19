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

	n, m := nextInt(), nextInt()
	dp := make([]int, n+1)
	plan := make([]int, m)
	for i := 0; i < m; i++ {
		plan[i] = nextInt()
	}

	price := make([][3]int, n)
	for i := 1; i < n; i++ {
		price[i][0] = nextInt()
		price[i][1] = nextInt()
		price[i][2] = nextInt()
	}

	for i := 0; i < m-1; i++ {
		start := min(plan[i], plan[i+1])
		end := max(plan[i], plan[i+1])
		dp[start]++
		dp[end]--
	}

	visit := 0
	result := 0
	for i := 1; i < n; i++ {
		visit += dp[i]
		result += min(price[i][0]*visit, price[i][1]*visit+price[i][2])
	}

	fmt.Fprintln(wr, result)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
