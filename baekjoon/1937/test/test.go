package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	//r  = bufio.NewReader(os.Stdin)
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	n, ans         int
	graph, visited [501][501]int
	dx, dy         []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n = nextInt()

	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			graph[i][j] = nextInt()
		}
	}

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(x, y int) int {
	if visited[x][y] == 0 {
		visited[x][y] = 1

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx > 0 && ny > 0 && nx <= n && ny <= n {
				if graph[x][y] < graph[nx][ny] {
					visited[x][y] = max(visited[x][y], dfs(nx, ny)+1)
				}
			}
		}
	}

	return visited[x][y]
}
