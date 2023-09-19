package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc        = bufio.NewScanner(os.Stdin)
	wr        = bufio.NewWriter(os.Stdout)
	n, m      int
	water     [][]int
	move      [][]int
	graph     [][]byte
	moveGraph [][]byte
	start     [2]int
	end       [2]int
	movq      [][2]int
	dx        = []int{-1, 0, 1, 0}
	dy        = []int{0, 1, 0, -1}
	q         [][2]int
)

func nextWord() []byte {
	sc.Scan()
	return sc.Bytes()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	water = make([][]int, n)
	move = make([][]int, n)
	for i := 0; i < n; i++ {
		water[i] = make([]int, m)
		move[i] = make([]int, m)
	}

	graph = make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		graph = append(graph, nextWord())
		for j, v := range graph[i] {
			if v == '*' {
				q = append(q, [2]int{i, j})
			} else if v == 'S' {
				start = [2]int{i, j}
			} else if v == 'D' {
				end = [2]int{i, j}
			}
		}
	}

	moveGraph = make([][]byte, n)
	for i := 0; i < n; i++ {
		moveGraph[i] = make([]byte, m)
		copy(moveGraph[i], graph[i])
	}

	dfsWater()
	water[end[0]][end[1]] = 100000
	dfsMove()

	//for _, v := range move {
	//	fmt.Println(v)
	//}

	if move[end[0]][end[1]] == 0 {
		fmt.Fprintln(wr, "KAKTUS")
		return
	} else {
		fmt.Fprintln(wr, move[end[0]][end[1]])
	}

}

func dfsWater() {
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if graph[nx][ny] == '.' {
				q = append(q, [2]int{nx, ny})
				graph[nx][ny] = '*'
				water[nx][ny] = water[x][y] + 1
			}
		}
	}
}

func dfsMove() {
	movq = make([][2]int, 0, n*m)
	movq = append(movq, start)
	for len(movq) > 0 {
		x, y := movq[0][0], movq[0][1]
		if x == end[0] && y == end[1] {
			break
		}
		movq = movq[1:]
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if moveGraph[nx][ny] == '.' && (water[nx][ny] > move[x][y]+1 || water[nx][ny] == 0) {
				movq = append(movq, [2]int{nx, ny})
				moveGraph[nx][ny] = 'S'
				move[nx][ny] = move[x][y] + 1
			} else if moveGraph[nx][ny] == 'D' {
				movq = append(movq, [2]int{nx, ny})
				moveGraph[nx][ny] = 'S'
				move[nx][ny] = move[x][y] + 1
			}
		}
	}

}
