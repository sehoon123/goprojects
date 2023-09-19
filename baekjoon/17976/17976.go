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
	n := nextInt()
	knots := make([][2]int, n)
	for i := 0; i < n; i++ {
		start, end := nextInt(), nextInt()
		knots[i] = [2]int{start, start + end}
	}

	sort.Slice(knots, func(i, j int) bool {
		return knots[i][0] < knots[j][0]
	})

	result := binary(1, 2000000001, knots)
	fmt.Fprintln(wr, result)

}

func binary(left, right int, knots [][2]int) int {
	for left <= right {
		mid := (left + right) / 2
		if check(mid, knots) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func check(mid int, knots [][2]int) bool {
	start := knots[0][0]
	for i := 0; i < len(knots); i++ {
		if start >= knots[i][0] && start <= knots[i][1] {
			start += mid
		} else if start < knots[i][0] {
			start = knots[i][0] + mid
		} else {
			return false
		}
	}
	return true
}
