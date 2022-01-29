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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()
	k := nextInt()
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	temp := n * n
	index_x := 0
	index_y := 0
	dy := []int{0, 1, 0, -1}
	dx := []int{1, 0, -1, 0}
	x, y := 0, 0
	graph[0][0] = temp
	temp--
	i := 0
	for {
		nx := x + dx[i%4]
		ny := y + dy[i%4]
		if nx >= n || ny >= n || nx < 0 || ny < 0 {
			i++
			continue
		}
		if graph[nx][ny] != 0 {
			i++
			continue
		}
		x = nx
		y = ny
		graph[x][y] = temp
		if temp == k {
			index_x = x
			index_y = y
		}
		temp--
		if temp == 0 {
			break
		}
	}

	for _, v := range graph {
		for _, vv := range v {
			fmt.Fprintf(wr, "%d ", vv)
		}
		fmt.Fprintln(wr)
	}

	fmt.Fprintln(wr, index_x+1, index_y+1)
}
