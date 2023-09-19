package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	q       [][]int
	graph   []string
	visited [][][]int
	n, m, k int
	dx      []int
	dy      []int
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &m, &k)
	q = make([][]int, 0, n*m)
	visited = make([][][]int, n)
	graph = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &graph[i])
		visited[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([]int, k+1)
		}
	}
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}

	fmt.Fprintln(w, bfs())

}

func bfs() int {
	q = append(q, []int{0, 0, k})
	visited[0][0][k] = 1
	for len(q) > 0 {
		x, y, z := q[0][0], q[0][1], q[0][2]
		if x == n-1 && y == m-1 {
			return visited[x][y][z]
		}
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if graph[nx][ny] == '1' && z > 0 && visited[nx][ny][z-1] == 0 {
				visited[nx][ny][z-1] = visited[x][y][z] + 1
				q = append(q, []int{nx, ny, z - 1})
			} else if graph[nx][ny] == '0' && visited[nx][ny][z] == 0 {
				visited[nx][ny][z] = visited[x][y][z] + 1
				q = append(q, []int{nx, ny, z})
			}
		}

	}
	return -1
}
