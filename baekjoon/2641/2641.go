package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n := nextInt()
	sample := make([]int, n)
	for i := 0; i < n; i++ {
		sample[i] = nextInt()
	}
	reverse := make([]int, n)
	for i := 0; i < n; i++ {
		switch sample[i] {
		case 1:
			reverse[n-1-i] = 3
		case 2:
			reverse[n-1-i] = 4
		case 3:
			reverse[n-1-i] = 1
		case 4:
			reverse[n-1-i] = 2
		}
	}

	m := nextInt()
	result := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		test := make([]int, n)
		for j := 0; j < n; j++ {
			test[j] = nextInt()
		}
		for j := 1; j <= n; j++ {
			tmp := rotate(test, j)
			if reflect.DeepEqual(tmp, reverse) || reflect.DeepEqual(tmp, sample) {
				result = append(result, test)
			}
		}
	}
	fmt.Fprintln(wr, len(result))
	for _, v := range result {
		for _, vv := range v {
			fmt.Fprint(wr, vv, " ")
		}
		fmt.Fprintln(wr)
	}
}

func rotate(arr []int, k int) []int {
	return append(arr[k:], arr[:k]...)
}
