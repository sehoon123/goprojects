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

	var n, c int
	fmt.Fscan(r, &n, &c)

	houses := make([]int, n)
	for i := range houses {
		fmt.Fscan(r, &houses[i])
	}

	wifi := make([]int, 0, c)

	sort.Ints(houses)
	start, end := 1, houses[n-1]

	for start < end {
		wifi = make([]int, 0, c)
		wifi = append(wifi, houses[0])
		mid := (start + end) / 2
		for _, v := range houses {
			if v-wifi[len(wifi)-1] >= mid {
				wifi = append(wifi, v)
			}
		}
		if len(wifi) >= c {
			start = mid + 1
		} else {
			end = mid
		}

	}
	fmt.Fprintln(w, end-1)

}
