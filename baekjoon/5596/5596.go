package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sum_a := 0
	sum_b := 0
	for i := 0; i < 4; i++ {
		sum_a += nextInt()
	}
	for i := 0; i < 4; i++ {
		sum_b += nextInt()
	}

	fmt.Fprintln(wr, max(sum_b, sum_a))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
