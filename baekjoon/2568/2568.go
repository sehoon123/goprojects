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
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		arr[i][0], arr[i][1] = nextInt(), nextInt()
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	l := 0
	for i := 0; i < n; i++ {
		lo, hi := 1, l+1
		for lo < hi {
			mid := (lo + hi) / 2
			if arr[m[mid]][1] < arr[i][1] {
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

	fmt.Fprintln(wr, n-len(s))
	//for i := 0; i < l; i++ {
	//	fmt.Fprintf(wr, "%d ", arr[s[i]][0])
	//}

	for i := 0; i < n; i++ {
		if !contains(s, i) {
			fmt.Fprintf(wr, "%d\n", arr[i][0])
		}
	}
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
