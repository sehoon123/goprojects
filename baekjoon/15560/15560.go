package main

import (
	"bufio"
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

	n, q, u, v := nextInt(), nextInt(), nextInt(), nextInt()

	arr := make([]int, n)
	accArr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
		if i == 0 {
			accArr[i] = arr[i]
		} else {
			accArr[i] = accArr[i-1] + arr[i]
		}
	}

}
