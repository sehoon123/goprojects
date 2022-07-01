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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	l, r := 0, n-1
	ans := abs(arr[l] + arr[r])
	x, y := arr[l], arr[r]

	for l < r {
		tmp := arr[l] + arr[r]
		if abs(tmp) < abs(ans) {
			ans = abs(tmp)
			x, y = arr[l], arr[r]
		}

		if tmp < 0 {
			l++
		} else {
			r--
		}
	}

	fmt.Fprintln(wr, x, y)

}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
