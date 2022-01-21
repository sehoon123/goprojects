package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	distance := make([]int, n-1)
	for i := range distance {
		fmt.Fscan(r, &distance[i])
	}

	cheap := 1000000001
	price := make([]int, n)
	for i := range price {
		fmt.Fscan(r, &price[i])
	}

	total := 0
	for i := 0; i < n-1; i++ {
		if price[i] < cheap {
			cheap = price[i]
		}
		total += distance[i] * cheap
	}

	fmt.Fprintln(w, total)
}
