package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc      = bufio.NewScanner(os.Stdin)
	wr      = bufio.NewWriter(os.Stdout)
	graph   [501][501]int
	visited [501][501]int
	dx, dy  []int
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

	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			graph[i][j] = nextInt()
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res := dfs(i, j)
			if ans < res {
				ans = res
			}
		}
	}

	fmt.Fprintln(wr, ans)
}

func dfs(x, y int) int {
	if visited[x][y] == 0 {
		visited[x][y] = 1
	}
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph) {
			continue
		}
		if graph[x][y] < graph[nx][ny] {
			visited[x][y] = max(visited[x][y], dfs(nx, ny)+1)
		}
	}
	return visited[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
