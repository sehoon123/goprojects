package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	wr  = bufio.NewWriter(os.Stdout)
	arr []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m, e := nextInt(), nextInt(), nextInt()
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	left, right := e, e
	count := 0
	for count < m {
		//lidx := binary(left)
		lidx := sort.Search(n, func(i int) bool {
			return arr[i] >= left
		})
		ridx := sort.Search(n, func(i int) bool {
			return arr[i] > right
		})
		if lidx != 0 {
			lidx--
		} else {
			right = arr[ridx]
			count++
			continue
		}
		if arr[ridx]-right < left-arr[lidx] {
			right = arr[ridx]
			count++
		} else {
			left = arr[lidx]
			count++
		}
	}

	fmt.Fprintln(wr, right-left)

}

func binary(n int) int {
	left, right := 0, len(arr)
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
