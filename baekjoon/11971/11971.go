package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)

	limit := make([]int, 0, 100)
	speed := make([]int, 0, 100)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		for j := 0; j < a; j++ {
			limit = append(limit, b)
		}
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		for j := 0; j < a; j++ {
			speed = append(speed, b)
		}
	}

	max := 0
	for i := 0; i < 100; i++ {
		if speed[i]-limit[i] > max {
			max = speed[i] - limit[i]
		}
	}

	fmt.Fprintln(w, max)
}
