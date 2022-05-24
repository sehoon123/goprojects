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
	graph  [][]int
	dx, dy []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextChar() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	m := nextInt()

	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	var apple [][]int
	var time []int
	var direction []string
	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}

	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		apple = append(apple, []int{a, b})
	}

	for i := 0; i < m; i++ {
		graph[apple[i][0]-1][apple[i][1]-1] = 1
	}

	l := nextInt()

	for i := 0; i < l; i++ {
		a, b := nextInt(), nextChar()
		time = append(time, a)
		direction = append(direction, b)
	}

	fmt.Fprintln(wr, graph)

	x, y := 0, 0
	for {
		if x < 0 || x >= n || y < 0 || y >= n {
			break
		}

	}
}
