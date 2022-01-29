package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	count int
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
	n, s := nextInt(), nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	solve(0, 0, s, a)
	wr.WriteString(strconv.Itoa(count) + "\n")

}

func solve(sum, i, s int, a []int) {
	if i == len(a) {
		return
	}
	if sum+a[i] == s {
		count++
	}
	// Include a[i]
	solve(sum+a[i], i+1, s, a)
	// Exclude a[i]
	solve(sum, i+1, s, a)
}
