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

	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	result := develop(0, n-1, a)
	fmt.Fprintln(wr, result)

}

func develop(start, end int, a []int) int {
	if start+1 == end {
		return max(a[start], a[end])
	}
	mid := (start + end) / 2
	LM, RM := 0, 0
	for i := start; i <= mid; i++ {
		LM = max(LM, a[i])
	}
	for i := mid + 1; i <= end; i++ {
		RM = max(RM, a[i])
	}
	return max(develop(start, mid, a)+RM, develop(mid+1, end, a)+LM)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
