package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)

	stuff := make([][2]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscanf(r, "%d %d\n", &stuff[i][0], &stuff[i][1])
	}

	d := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([]int, k+1)
	}

	//stuff[i][0] : weight, stuff[i][1] : value
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j >= stuff[i][0] {
				d[i][j] = max(d[i-1][j], d[i-1][j-stuff[i][0]]+stuff[i][1])
			} else {
				d[i][j] = d[i-1][j]
			}
		}
	}

	fmt.Fprintln(w, d[n][k])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
