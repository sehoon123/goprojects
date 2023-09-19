package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, wr := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	var N, M int
	fmt.Fscan(r, &N, &M)

	var INF = 1000000000

	graph := make([][]int, 101)
	for i := 0; i < 101; i++ {
		graph[i] = make([]int, 101)
		for j := range graph[i] {
			graph[i][j] = INF
		}
	}

	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(r, &u, &v)
		graph[u][v] = 1
		graph[v][u] = 1
	}

	for k := 1; k < N+1; k++ {
		for a := 1; a < N+1; a++ {
			for b := 1; b < N+1; b++ {
				graph[a][b] = min(graph[a][b], graph[a][k]+graph[k][b])
			}
		}
	}

	var X, K int
	fmt.Fscan(r, &X, &K)

	distance := graph[1][K] + graph[K][X]

	if distance >= INF {
		fmt.Fprintln(wr, -1)
	} else {
		fmt.Fprintln(wr, distance)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
