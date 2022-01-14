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
	var n, m int
	defer w.Flush()
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	array := []int{}
	visited := make([]bool, n+1)

	//for i := 1; i <= n; i++ {
	//	var a int
	//	fmt.Fscan(r, &a)
	//	array = append(array, a)
	//}

	dfs(array, n, m, &visited)

}

func dfs(array []int, n, m int, visited *[]bool) {
	if len(array) == m {
		for _, v := range array {
			fmt.Fprintf(w, "%d ", v)
		}
		fmt.Fprintln(w)
		return
	}

	for i := 1; i <= n; i++ {
		if (*visited)[i] {
			continue
		}
		(*visited)[i] = true
		array = append(array, i)
		dfs(array, n, m, visited)
		array = array[:len(array)-1]
		(*visited)[i] = false
	}

}
