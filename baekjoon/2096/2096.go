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
	graph := make([][3]int, n)
	dp_max := make([][3]int, n)
	dp_min := make([][3]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			graph[i][j] = nextInt()
		}
	}

	for i := 0; i < 3; i++ {
		dp_max[0][i] = graph[0][i]
		dp_min[0][i] = graph[0][i]
	}

	for i := 1; i < n; i++ {
		dp_max[i][0] = max(dp_max[i-1][0], dp_max[i-1][1]) + graph[i][0]
		dp_max[i][1] = max(max(dp_max[i-1][0], dp_max[i-1][1]), dp_max[i-1][2]) + graph[i][1]
		dp_max[i][2] = max(dp_max[i-1][1], dp_max[i-1][2]) + graph[i][2]
		dp_min[i][0] = min(dp_min[i-1][0], dp_min[i-1][1]) + graph[i][0]
		dp_min[i][1] = min(min(dp_min[i-1][0], dp_min[i-1][1]), dp_min[i-1][2]) + graph[i][1]
		dp_min[i][2] = min(dp_min[i-1][1], dp_min[i-1][2]) + graph[i][2]
	}

	fmt.Fprintf(wr, "%v %v", max(max(dp_max[n-1][0], dp_max[n-1][1]), dp_max[n-1][2]), min(min(dp_min[n-1][0], dp_min[n-1][1]), dp_min[n-1][2]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
