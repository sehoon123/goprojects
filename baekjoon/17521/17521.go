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

	n, w := nextInt(), nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			coin := w / arr[i-1]
			w -= coin * arr[i-1]
			w += coin * arr[i]
		}
	}

	fmt.Fprintln(wr, w)
}
