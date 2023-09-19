package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	T := nextInt()

	n := nextInt()
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	m := nextInt()
	B := make([]int, m)
	for i := 0; i < m; i++ {
		B[i] = nextInt()
	}

	var accA, accB []int

	for i := 0; i < n; i++ {
		sum := A[i]
		accA = append(accA, sum)
		for j := i + 1; j < n; j++ {
			sum += A[j]
			accA = append(accA, sum)
		}
	}

	for i := 0; i < m; i++ {
		sum := B[i]
		accB = append(accB, sum)
		for j := i + 1; j < m; j++ {
			sum += B[j]
			accB = append(accB, sum)
		}
	}

	sort.Ints(accB)

	ans := 0

	for _, v := range accA {
		diff := T - v
		upper := upperBound(accB, diff)
		lower := lowerBound(accB, diff)

		ans += upper - lower
	}

	fmt.Fprintln(wr, ans)
}

func upperBound(arr []int, target int) int {
	low, high, mid := 0, len(arr)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func lowerBound(arr []int, target int) int {
	low, high, mid := 0, len(arr)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
