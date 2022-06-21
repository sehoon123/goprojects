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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, k := nextInt(), nextInt()
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nextInt()
	}

	max := -987654321
	for i := k; i <= n; i++ {
		if arr[i]-arr[i-k] > max {
			max = arr[i] - arr[i-k]
		}
	}

	//fmt.Fprintln(wr, arr)
	fmt.Fprintln(wr, max)
}
