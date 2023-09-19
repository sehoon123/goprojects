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
	graph   [][]int
	virus   [][]int
	visited [][]int
	dx      []int
	dy      []int
	q       [][]int
	n, m    int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m = nextInt(), nextInt()
	dx = []int{-1, 0, 0, 1}
	dy = []int{0, 1, -1, 0}
	virus = make([][]int, 0, m)
	graph = make([][]int, n)
	visited = make([][]int, n)
	q = make([][]int, 0, n*n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		visited[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = nextInt()
			if graph[i][j] == 2 {
				virus = append(virus, []int{i, j})
			}
		}
	}

	fmt.Fprintln(wr, graph)
	fmt.Fprintln(wr, virus)
	q = append(q, []int{0, 0})
	bfs()
	fmt.Fprintln(wr, visited)

}

func bfs() {
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				continue
			}
			if visited[nx][ny] == 0 && (graph[nx][ny] == 0 || graph[nx][ny] == 2) {
				visited[nx][ny] = visited[x][y] + 1
				q = append(q, []int{nx, ny})
			}
		}
	}
}
