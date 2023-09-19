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

	n := nextInt()

	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 1
	arr[2] = 2

	for i := 3; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	fmt.Fprintln(wr, arr[n-1], n-2)
}
