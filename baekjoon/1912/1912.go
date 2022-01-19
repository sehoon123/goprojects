package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	array := make([]int, n)
	for i, v := range array {
		fmt.Fscanf(r, "%d", &v)
		array[i] = v
	}

	d := make([]int, n)
	d[0] = array[0]
	for i := 1; i < n; i++ {
		d[i] = max(array[i], d[i-1]+array[i])
	}
	sort.Ints(d)
	fmt.Fprintln(w, d[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
