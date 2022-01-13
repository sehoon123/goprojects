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
	A := []int{}
	B := []int{}

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		A = append(A, a)
	}
	for i := 0; i < n; i++ {
		var b int
		fmt.Fscan(r, &b)
		B = append(B, b)
	}

	sort.Ints(A)
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	for i := 0; i < k; i++ {
		if A[i] < B[i] {
			A[i], B[i] = B[i], A[i]
		}
	}
	sum := 0
	for _, v := range A {
		sum += v
	}

	fmt.Println(sum)
}
