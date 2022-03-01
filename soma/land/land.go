package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	arr   []int
	cache []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	cache = make([]int, n+1)

	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}
	result := solve(0, n-1)
	fmt.Println(result)
}

func solve(start, end int) int {
	if start+1 == end {
		return max(arr[start], arr[end])
	}
	mid := (start + end) / 2
	leftMax, rightMax := 0, 0

	for i := start; i < mid+1; i++ {
		leftMax = max(leftMax, arr[i])
	}
	for i := mid + 1; i < end; i++ {
		rightMax = max(rightMax, arr[i])
	}

	return max(solve(start, mid)+rightMax, solve(mid+1, end)+leftMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
