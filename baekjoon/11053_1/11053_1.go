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
	var n int
	fmt.Fscan(r, &n)

	seq := make([]int, n)
	tmp := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &seq[i])
	}

	for _, v := range seq {
		push(&tmp, v)
	}

	fmt.Fprintf(w, "%d\n", len(tmp))
}

func push(tmp *[]int, v int) {
	if len(*tmp) == 0 {
		*tmp = append(*tmp, v)
		return
	}

	if v > (*tmp)[len(*tmp)-1] {
		*tmp = append(*tmp, v)
		return
	}

	if v < (*tmp)[len(*tmp)-1] {
		k := LowerBound(*tmp, v)
		(*tmp)[k] = v
		return
	}

}

func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
