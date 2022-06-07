package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	wr         = bufio.NewWriter(os.Stdout)
	graph      [5][5][5]int
	visit      [5][5][5]int
	dx, dy, dz []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	dx = []int{1, 0, -1, 0, 0, 0}
	dy = []int{0, 1, 0, -1, 0, 0}
	dz = []int{0, 0, 0, 0, 1, -1}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				graph[i][j][k] = nextInt()
			}
		}
	}

	bfs(0, 0, 0)
	fmt.Fprintln(wr, "----------------------------------------")
	for _, v := range visit {
		for _, vv := range v {
			fmt.Fprintln(wr, vv)
		}
		fmt.Fprintln(wr)
	}
}

func bfs(x, y, z int) {
	queue := make([]int, 0)
	queue = append(queue, x, y, z)
	visit[x][y][z] = 1
	for len(queue) > 0 {
		x, y, z = queue[0], queue[1], queue[2]
		queue = queue[3:]
		for i := 0; i < 6; i++ {
			nx, ny, nz := x+dx[i], y+dy[i], z+dz[i]
			if nx < 0 || nx >= 5 || ny < 0 || ny >= 5 || nz < 0 || nz >= 5 {
				continue
			}
			if graph[nx][ny][nz] == 1 && visit[nx][ny][nz] == 0 {
				queue = append(queue, nx, ny, nz)
				visit[nx][ny][nz] = visit[x][y][z] + 1
			}
		}
	}
}
