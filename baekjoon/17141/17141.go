package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc      = bufio.NewScanner(os.Stdin)
	wr      = bufio.NewWriter(os.Stdout)
	graph   [][]int
	dx, dy  []int
	visited [][]int
	q       [][]int
	success []int
	temp    []int
	max     int
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
	m := nextInt()

	graph = make([][]int, n)
	visited = make([][]int, n)
	virus := make([][]int, 0, 50)
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}

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

	success = make([]int, len(virus))

	result := make([][]int, 0)
	temp = make([]int, 0)

	combination(virus, result, m)

	if len(temp) == 0 {
		wr.WriteString("-1\n")
		return
	}
	sort.Ints(temp)

	fmt.Fprintln(wr, temp[0])

}

func bfs() {
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph) {
				continue
			}
			if (visited[nx][ny] == 0 || visited[nx][ny] >= visited[x][y]+1) && (graph[nx][ny] == 0 || graph[nx][ny] == 2) {
				visited[nx][ny] = visited[x][y] + 1
				q = append(q, []int{nx, ny})
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			if (graph[i][j] == 2 || graph[i][j] == 0) && visited[i][j] == 0 {
				return
			}
			if visited[i][j] > max {
				max = visited[i][j]
			}
		}
	}
	temp = append(temp, max-1)
}

func combination(arr, result [][]int, m int) {
	if len(result) == m {
		visited = make([][]int, len(graph))
		for i := 0; i < len(graph); i++ {
			visited[i] = make([]int, len(graph))
		}

		for i := 0; i < m; i++ {
			q = append(q, []int{result[i][0], result[i][1]})
			visited[result[i][0]][result[i][1]] = 1
		}
		max = -1
		bfs()
		return
	}

	for i := 0; i < len(arr); i++ {
		if success[i] == 0 {
			success[i] = 1
			result = append(result, arr[i])
			combination(arr, result, m)
			result = result[:len(result)-1]
			for j := i + 1; j < len(arr); j++ {
				success[j] = 0
			}
		}
	}
}
