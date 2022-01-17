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
	fmt.Fscan(r, &n)

	seq := make([]int, n)
	d := make([]int, n)
	var max int

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &seq[i])
		for j := 0; j < i; j++ {
			if seq[i] > seq[j] && d[i] < d[j] {
				d[i] = d[j]
			}
		}
		d[i]++
		if d[i] > max {
			max = d[i]
		}
	}

	fmt.Fprintln(w, max)

}
