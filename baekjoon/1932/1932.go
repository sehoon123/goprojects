package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Fscan(r, &graph[i][j])
		}
	}

	var left, right int
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				left = 0
			} else {
				left = graph[i-1][j-1]
			}

			if j > i-1 {
				right = 0
			} else {
				right = graph[i-1][j]
			}

			graph[i][j] += int(math.Max(float64(left), float64(right)))

		}
	}

	sort.Ints(graph[n-1][:])
	fmt.Fprintln(w, graph[n-1][n-1])

}
