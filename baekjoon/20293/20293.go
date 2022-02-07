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

	INF := int(1e9)

	R, C := nextInt(), nextInt()
	N := nextInt()
	graph := [3002][3002]int{}

	gasStation := make([][3]int, N)
	for i := 0; i < N; i++ {
		gasStation[i][0], gasStation[i][1], gasStation[i][2] = nextInt(), nextInt(), nextInt()
		graph[gasStation[i][0]][gasStation[i][1]] = gasStation[i][2]
	}

	for i := R; i > 0; i-- {
		graph[i][C+1] = INF

		for j := C; j > 0; j-- {
			graph[R+1][j] = INF

			if i == R && j == C {
				graph[i][j] = 0
				continue
			}

			t := min(graph[i+1][j], graph[i][j+1]) + 1
			if graph[i][j] > 0 {
				t -= graph[i][j]
				if t < 0 {
					t = 0
				}
			}

			graph[i][j] = t

		}
	}

	fmt.Fprintln(wr, graph[1][1])

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
