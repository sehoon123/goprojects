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
	N := nextInt()
	arr := make([]int, 0, N)
	for i := 0; i < N; i++ {
		arr = Insert(arr, nextInt())
		fmt.Fprintln(wr, arr[i/2])
	}
}

func Insert(arr []int, x int) []int {
	i := sort.Search(len(arr), func(i int) bool { return arr[i] > x })
	arr = append(arr, 0)
	copy(arr[i+1:], arr[i:])
	arr[i] = x
	return arr
}
