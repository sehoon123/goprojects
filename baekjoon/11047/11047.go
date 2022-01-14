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
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)

	money_list := []int{}

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		money_list = append(money_list, a)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(money_list)))

	count := 0
	for k > 0 {
		for _, v := range money_list {
			if k >= v {
				count += k / v
				k = k - (v * (int(k / v)))
			}
		}
	}

	fmt.Fprintln(w, count)
}
