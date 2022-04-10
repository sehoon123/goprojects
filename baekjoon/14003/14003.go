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
	p := make([]int, n)
	m := make([]int, n+1)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	l := 0
	for i := 0; i < n; i++ {
		lo, hi := 1, l+1
		for lo < hi {
			mid := (lo + hi) / 2
			if arr[m[mid]] < arr[i] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		newl := lo

		p[i] = m[newl-1]
		m[newl] = i
		if newl > l {
			l = newl
		}
	}

	s := make([]int, l)
	k := m[l]
	for i := l - 1; i >= 0; i-- {
		s[i] = k
		k = p[k]
	}

	fmt.Fprintln(wr, len(s))
	for i := 0; i < l; i++ {
		fmt.Fprintf(wr, "%d ", arr[s[i]])
	}
}
