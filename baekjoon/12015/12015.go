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

	array := make([]int, n)
	lis := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &array[i])
	}

	lis = append(lis, array[0])
	for i := 1; i < n; i++ {
		if array[i] > lis[len(lis)-1] {
			lis = append(lis, array[i])
		} else {
			k := upperbound(lis, array[i])
			lis[k] = array[i]
		}
	}

	fmt.Fprintln(w, len(lis))

}

func upperbound(array []int, target int) int {
	start, end := 0, len(array)-1
	for start < end {
		mid := (start + end) / 2
		if array[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
