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
	for _, v := range sc.Bytes() {
		if v == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(v - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	INF := int(1e9)
	n := nextInt()
	m := nextInt()

	graph := make([][]int, 101)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, 101)
		for j := range graph[i] {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = INF
			}
		}
	}

	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		if graph[a][b] > c {
			graph[a][b] = c
		}
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if graph[i][j] >= INF || i == j {
				fmt.Fprintf(wr, "%d ", 0)
			} else {
				fmt.Fprintf(wr, "%d ", graph[i][j])
			}
		}
		fmt.Fprintln(wr)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
