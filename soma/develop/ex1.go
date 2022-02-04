package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, v := range sc.Bytes() {
		if v == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(v - '0')
	}

	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	//total := 0
	//result := develop(0, n, total, a)
	result := develop(0, n, a)

	fmt.Fprintln(wr, result)

}

//func develop(start, end, total int, a []int) int {
//	if len(a) == 2 {
//		total += max(a)
//		return total
//	}
//	mid := (start + end) / 2
//	if max(a[:mid]) > max(a[mid:]) {
//		total += max(a[mid:])
//		return develop(start, mid, total, a[:mid])
//	} else {
//		total += max(a[:mid])
//		return develop(mid, end, total, a[mid:])
//	}
//}

func develop(start, end int, a []int) int {
	if end-start == 1 {
		return max(a[start:end])
	}
	mid := (start + end) / 2
	leftMax := max(a[start:mid])
	rightMax := max(a[mid:end])
	left := develop(start, mid, a) + rightMax
	right := develop(mid, end, a) + leftMax
	return max1(left, right)
}

func max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func max1(a, b int) int {
	if a < b {
		return b
	}
	return a
}
