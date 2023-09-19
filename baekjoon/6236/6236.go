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
var n int
var m int
var spend []int

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

	n, m = nextInt(), nextInt()
	spend = make([]int, n)
	for i := 0; i < n; i++ {
		spend[i] = nextInt()
	}

	left, right := 0, int(1e9)
	result := binary(left, right)
	fmt.Fprintln(wr, result)

}

func binary(left, right int) int {
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

func check(mid int) bool {
	now := 0
	cnt := 0
	for i := 0; i < n; i++ {
		if mid < spend[i] {
			return false
		}
		if now < spend[i] {
			now = mid - spend[i]
			cnt++
		} else {
			now -= spend[i]
		}
	}
	return cnt <= m
}
