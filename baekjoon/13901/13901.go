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

const INF = int(1e9)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	r, c := nextInt(), nextInt()
	k := nextInt()

	visited := make([][]int, r)
	for i := range visited {
		visited[i] = make([]int, c)
	}

	// 장애물 방문처리
	for i := 0; i < k; i++ {
		x, y := nextInt(), nextInt()
		visited[x][y] = 1
	}

	sx, sy := nextInt(), nextInt()
	// 시작지점 방문처리
	visited[sx][sy] = 1
	direction := make([]int, 4)
	for i := 0; i < 4; i++ {
		direction[i] = nextInt()
	}

	i := 0
	count := 0
	for {
		nx, ny := 0, 0
		if direction[i] == 1 {
			nx, ny = sx-1, sy
		} else if direction[i] == 2 {
			nx, ny = sx+1, sy
		} else if direction[i] == 3 {
			nx, ny = sx, sy-1
		} else {
			nx, ny = sx, sy+1
		}

		if nx < 0 || nx >= r || ny < 0 || ny >= c {
			i++
			i %= 4
			count++
			continue
		}

		if visited[nx][ny] == 0 {
			visited[nx][ny] = 1
			sx, sy = nx, ny
			count = 0
		} else {
			i++
			i %= 4
			count++
		}

		if count == 4 {
			fmt.Fprintln(wr, sx, sy)
			return
		}

	}
}
