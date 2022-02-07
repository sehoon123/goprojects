package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

type queue struct {
	data [][2]int
}

func (q *queue) push(x, y int) {
	q.data = append(q.data, [2]int{x, y})
}

func (q *queue) pop() (x, y int) {
	x, y = q.data[0][0], q.data[0][1]
	q.data = q.data[1:]
	return
}

func (q *queue) empty() bool {
	return len(q.data) == 0
}

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	r := bufio.NewReader(os.Stdin)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, w := nextInt(), nextInt()

	pond := make([][2]int, w)
	for i := 0; i < w; i++ {
		pond[i][0] = nextInt() - 1
		pond[i][1] = nextInt() - 1
	}

	graph := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &graph[i])
	}

	water := make([][]int, n)
	for i := 0; i < n; i++ {
		water[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			water[i][j] = -1
		}
	}

	q := queue{}
	for i := 0; i < w; i++ {
		q.push(pond[i][0], pond[i][1])
		water[pond[i][0]][pond[i][1]] = 0
	}

	left, right := 0, WaterDay(water, q)

	answer := binary(left, right, graph, water)
	fmt.Fprintln(wr, answer)

}

func WaterDay(water [][]int, q queue) int {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	maxWater := 0
	for !q.empty() {
		x, y := q.pop()
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if 0 <= nx && nx < len(water) && 0 <= ny && ny < len(water) {
				if water[nx][ny] == -1 {
					water[nx][ny] = water[x][y] + 1
					q.push(nx, ny)
				}
			}
		}
	}
	for i := 0; i < len(water); i++ {
		for j := 0; j < len(water); j++ {
			if water[i][j] > maxWater {
				maxWater = water[i][j]
			}
		}
	}
	return maxWater
}

func binary(left, right int, graph []string, water [][]int) int {
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if check(mid, graph, water) {
			right = mid - 1
			result = mid
		} else {
			left = mid + 1
		}
	}
	return result
}

func check(mid int, graph []string, water [][]int) bool {
	dx := []int{1, 0, -1, 0, 1, -1, -1, 1}
	dy := []int{0, 1, 0, -1, 1, 1, -1, -1}

	n := len(water)
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}
	visited[0][0] = 1

	location := queue{}
	location.push(0, 0)

	for !location.empty() {
		x, y := location.pop()
		if (x == n-2 && y == n-2) || (x == n-2 && y == n-1) || (x == n-1 && y == n-2) {
			return true
		}

		for i := 0; i < 8; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < n {
				if graph[nx][ny] == '1' && water[nx][ny] <= mid && visited[nx][ny] == 0 {
					visited[nx][ny] = 1
					location.push(nx, ny)
				}
			}
		}
	}
	return false
}
