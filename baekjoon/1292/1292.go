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

	start, end := nextInt(), nextInt()
	a := make([]int, 0, 500500)
	for i := 1; i <= end; i++ {
		if len(a) > end {
			break
		}
		for j := 0; j < i; j++ {
			a = append(a, i)
		}
	}
	fmt.Fprintln(wr, sum(a[start-1:end]))

}

func sum(a []int) int {
	r := 0
	for _, v := range a {
		r += v
	}
	return r
}
