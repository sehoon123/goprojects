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
	for i := 0; i < T; i++ {
		M := nextInt()
		fmt.Fprintln(wr, (M+1)/2)
		arr := make([]int, 0, M)
		for j := 0; j < M; j++ {
			arr = Insert(arr, nextInt())
			if j%2 == 0 {
				fmt.Fprintf(wr, "%d ", arr[j/2])
			}
		}
		fmt.Fprintln(wr)
	}
}

func Insert(arr []int, x int) []int {
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] > x
	})
	arr = append(arr, 0)
	copy(arr[i+1:], arr[i:])
	arr[i] = x
	return arr
}
