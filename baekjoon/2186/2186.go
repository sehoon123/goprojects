package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	wr         = bufio.NewWriter(os.Stdout)
	graph      [][]uint8
	dp         [][][]int
	target     string
	dx, dy     []int
	n, m, k, l int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}

	n, m, k = nextInt(), nextInt(), nextInt()
	graph = make([][]uint8, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]uint8, m)
		temp := nextString()
		for j := 0; j < m; j++ {
			graph[i][j] = temp[j]
		}
	}

	target = nextString()
	l = len(target)
	dp = make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, l)
			for t := 0; t < l; t++ {
				dp[i][j][t] = -1
			}
		}
	}

	ans := 0
	var start [][2]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == target[0] {
				start = append(start, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(start); i++ {
		a, b := start[i][0], start[i][1]
		ans += dfs(a, b, 0)
	}

	wr.WriteString(strconv.Itoa(ans) + "\n")
}

func dfs(x, y, t int) int {
	if x < 0 || x >= n || y < 0 || y >= m || graph[x][y] != target[t] {
		return 0
	}
	if t == l-1 {
		return 1
	}
	if dp[x][y][t] != -1 {
		return dp[x][y][t]
	}
	dp[x][y][t] = 0

	for j := 1; j <= k; j++ {
		for i := 0; i < 4; i++ {
			dp[x][y][t] += dfs(x+j*dx[i], y+j*dy[i], t+1)
		}
	}
	return dp[x][y][t]
}
