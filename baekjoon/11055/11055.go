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
	cache []int
)

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = nextInt()
	}

	cache = make([]int, n)
	for i := range cache {
		cache[i] = arr[i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				cache[i] = max(cache[i], cache[j]+arr[i])
			}
		}
	}

	res := 0
	for i := range cache {
		res = max(res, cache[i])
	}

	fmt.Fprintln(wr, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
