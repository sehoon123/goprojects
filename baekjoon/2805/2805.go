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

	var n, m int
	fmt.Fscan(r, &n, &m)

	trees := make([]int, n)
	for i := range trees {
		fmt.Fscan(r, &trees[i])
	}

	sort.Ints(trees)

	start, end := 1, trees[n-1]

	for start < end {
		remain := 0
		mid := (start + end) / 2
		for _, v := range trees {
			if v < mid {
				continue
			}
			remain += v - mid
		}
		if remain >= m {
			start = mid + 1
		} else {
			end = mid
		}

	}

	sum := 0
	for _, v := range trees {
		if v < start {
			continue
		}
		sum += v - start
	}
	if sum < m {
		fmt.Fprintln(w, start-1)
	} else {
		fmt.Fprintln(w, start)
	}

}
