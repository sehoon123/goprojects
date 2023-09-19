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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			graph[i][j] = 1e9
		}
	}

	for i := 0; i < m; i++ {
		x, y, z := nextInt(), nextInt(), nextInt()
		if graph[x][y] > z {
			graph[x][y] = z
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}

	min := int(1e9)
	for i := 1; i <= n; i++ {
		if graph[i][i] < min {
			min = graph[i][i]
		}
	}

	if min >= int(1e9) {
		fmt.Fprint(wr, -1)
	} else {
		fmt.Fprintln(wr, min)
	}
}

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
