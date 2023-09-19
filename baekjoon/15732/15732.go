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
	n, k, d := nextInt(), nextInt(), nextInt()
	rules := make([][3]int, k)
	for i := 0; i < k; i++ {
		rules[i][0], rules[i][1], rules[i][2] = nextInt(), nextInt(), nextInt()
	}

	left, right := 0, n
	answer := 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid, d, rules) {
			right = mid - 1
			answer = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Fprintln(wr, answer)
}

func check(mid, d int, rules [][3]int) bool {
	sum := 0
	for _, rule := range rules {
		num := min(mid, rule[1])
		if num >= rule[0] {
			sum += (num-rule[0])/rule[2] + 1
		}
	}
	return sum >= d
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
