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
	fmt.Fscan(r, &n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &A[i])
	}
	sort.Ints(A)

	fmt.Fscan(r, &m)
	B := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &B[i])
	}

	for i := 0; i < m; i++ {
		binary_search(A, B[i])
	}
}

func binary_search(array []int, target int) {
	start := 0
	end := len(array) - 1
	for start <= end {
		mid := (start + end) / 2
		if array[mid] == target {
			fmt.Println(1)
			return
		} else if array[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	fmt.Println(0)
	return
}

//func nextInt() int {
//	sc.Scan()
//	r, f := 0, 1
//	for _, c := range sc.Bytes() {
//		if c == '-' {
//			f = -1
//			continue
//		}
//		r *= 10
//		r += int(c - '0')
//	}
//	return r * f
//}
