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

func nextChar() string {
	sc.Scan()
	return sc.Text()
}

func insort(a []int, x int) []int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})

	a = append(a, 0)
	copy(a[idx+1:], a[idx:])
	a[idx] = x
	return a
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	T := nextInt()
	for i := 0; i < T; i++ {
		N := nextInt()
		arr := make([]int, 0, N)
		for j := 0; j < N; j++ {
			command, num := nextChar(), nextInt()
			if command == "I" {
				arr = insort(arr, num)
			}

			if command == "D" {
				if num == 1 {
					if len(arr) > 0 {
						arr = arr[:len(arr)-1]
					}
				} else {
					if len(arr) > 0 {
						arr = arr[1:]
					}
				}
			}
		}

		if len(arr) == 0 {
			fmt.Fprintln(wr, "EMPTY")
		} else {
			fmt.Fprintln(wr, arr[len(arr)-1], arr[0])
		}
	}
}
