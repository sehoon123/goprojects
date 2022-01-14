package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	defer w.Flush()
	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	array := []int{}

	dfs(array, n, m)

}

func dfs(array []int, n, m int) {
	if len(array) == m {
		for _, v := range array {
			fmt.Fprintf(w, "%d ", v)
		}
		fmt.Fprintln(w)
		return
	}

	for i := 1; i <= n; i++ {
		array = append(array, i)
		dfs(array, n, m)
		array = array[:len(array)-1]
	}

}
