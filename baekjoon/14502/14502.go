package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	wr     = bufio.NewWriter(os.Stdout)
	n, m   int
	graph  [][]int
	temp   [][]int
	blanks [][]int
	bb     [][]int
	virus  [][]int
	vv     [][]int
	dx, dy []int
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
	graph = make([][]int, n)
	virus = make([][]int, 0, 100)
	blanks = make([][]int, 0, 100)
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
	max := -1

	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			graph[i][j] = nextInt()
			if graph[i][j] == 0 {
				blanks = append(blanks, []int{i, j})
			} else if graph[i][j] == 2 {
				virus = append(virus, []int{i, j})
			}
		}
	}

	for i := 0; i < len(blanks); i++ {
		for j := i + 1; j < len(blanks); j++ {
			for k := j + 1; k < len(blanks); k++ {
				now := bfs(i, j, k)
				if now > max {
					max = now
				}
			}
		}
	}

	wr.WriteString(strconv.Itoa(max) + "\n")
}

func bfs(a, b, c int) int {
	temp = make([][]int, n)
	vv = make([][]int, len(virus))
	for i := 0; i < n; i++ {
		temp[i] = make([]int, m)
		copy(temp[i], graph[i])
	}
	for i := 0; i < len(virus); i++ {
		vv[i] = make([]int, 2)
		copy(vv[i], virus[i])
	}

	temp[blanks[a][0]][blanks[a][1]] = 1
	temp[blanks[b][0]][blanks[b][1]] = 1
	temp[blanks[c][0]][blanks[c][1]] = 1
	for len(vv) > 0 {
		x, y := vv[0][0], vv[0][1]
		vv = vv[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if temp[nx][ny] == 0 {
				temp[nx][ny] = 2
				vv = append(vv, []int{nx, ny})
			}
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if temp[i][j] == 0 {
				cnt++
			}
		}
	}
	return cnt
}
