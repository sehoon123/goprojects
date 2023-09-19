package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n)

	parts := []int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		parts = append(parts, a)
	}

	fmt.Fscan(r, &m)
	demands := []int{}
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(r, &a)
		demands = append(demands, a)
	}

	for _, v := range demands {
		if binary_search(parts, v, 0, len(parts)-1) == -1 {
			fmt.Fprintln(w, "NO")
		} else {
			fmt.Fprintln(w, "YES")
		}

	}

}

func binary_search(array []int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if array[mid] == target {
		return mid
	} else if array[mid] > target {
		return binary_search(array, target, start, mid-1)
	} else {
		return binary_search(array, target, mid+1, end)
	}
}
