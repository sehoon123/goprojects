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

	arr := make([]int, 11)

	for i := 1; i <= 10; i++ {
		arr[i] = arr[i-1] + nextInt()
	}

	right, left := 10, 9
	for i := 1; i <= 10; i++ {
		if arr[i] > 100 {
			right = i
			left = i - 1
			break
		}
	}

	if abs(100-arr[right]) > abs(100-arr[left]) {
		fmt.Fprintln(wr, arr[left])
	} else {
		fmt.Fprintln(wr, arr[right])
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
