package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	wr     = bufio.NewWriter(os.Stdout)
	board  [][]int
	visit  [][]bool
	dx, dy []int
	count  int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m, k := nextInt(), nextInt(), nextInt()

	board = make([][]int, n)
	visit = make([][]bool, n)
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}

	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		visit[i] = make([]bool, m)
	}

	for i := 0; i < k; i++ {
		a, b := nextInt()-1, nextInt()-1
		board[a][b] = 1
	}

	max := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 1 {
				count = 1
				visit[i][j] = true
				bfs(i, j)
			}
			if max < count {
				max = count
			}
			count = 0
		}
	}

	fmt.Fprintln(wr, max)
}

func bfs(x, y int) {
	queue := make([]int, 0)
	queue = append(queue, x, y)

	for len(queue) > 0 {
		sx, sy := queue[0], queue[1]
		queue = queue[2:]
		for i := 0; i < 4; i++ {
			nx, ny := sx+dx[i], sy+dy[i]
			if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) {
				continue
			}
			if board[nx][ny] == 1 && !visit[nx][ny] {
				visit[nx][ny] = true
				queue = append(queue, nx, ny)
				count++
			}
		}
	}
}
