package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		arr[i][0], arr[i][1] = i, nextInt()
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

	result := -1
	for i := 0; i < n; i++ {
		if arr[i][0]-i > result {
			result = arr[i][0] - i
		}
	}

	fmt.Fprintln(wr, result+1)
}
