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
	fmt.Fscan(r, &n, &k)

	value := make([]int, n+1)
	weight := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(r, &weight[i], &value[i])
	}
	d := make([]int, k+1)
	fmt.Fprintln(w, d)

	for i := 1; i < n+1; i++ {
		for j := k; j >= weight[i]; j-- {
			d[j] = max(d[j], d[j-weight[i]]+value[i])
		}
	}

	for _, v := range d {
		fmt.Fprintln(w, v)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
