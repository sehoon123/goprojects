package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscan(r, &n, &m)

	temp := min(n, m)

	list := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &list[i])
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < temp; k++ {
				if i+k >= n || j+k >= m {
					continue
				}
				if list[i][j] == list[i+k][j] && list[i][j] == list[i+k][j+k] && list[i][j] == list[i][j+k] {
					result = max(result, k+1)
				}
			}
		}
	}

	fmt.Fprintln(w, result*result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
