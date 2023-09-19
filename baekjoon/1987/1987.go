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
	dx, dy  []int
	visited []int
	graph   []string
	r, c    int
	result  int
)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	r, c = nextInt(), nextInt()

	graph = make([]string, r)
	for i := 0; i < r; i++ {
		graph[i] = nextStr()
	}

	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
	visited = make([]int, 26)
	result = 0

	dfs(0, 0, 0)
	fmt.Fprintln(wr, result+1)
}

func dfs(x, y, count int) {
	if x < 0 || x >= r || y < 0 || y >= c {
		return
	}
	visited[graph[x][y]-'A'] = 1

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || nx >= r || ny < 0 || ny >= c {
			continue
		}
		if visited[graph[nx][ny]-'A'] == 0 {
			dfs(nx, ny, count+1)
		}
	}
	visited[graph[x][y]-'A'] = 0
	if count > result {
		result = count
	}
}
