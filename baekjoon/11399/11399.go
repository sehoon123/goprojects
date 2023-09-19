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

	list := make([]int, n)
	for i := range list {
		fmt.Fscanf(r, "%d", &list[i])
	}

	sort.Ints(list)
	sum := 0
	for i := range list {
		for j := 0; j <= i; j++ {
			sum += list[j]
		}
	}
	fmt.Fprintln(w, sum)
}
