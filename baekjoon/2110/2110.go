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

	n, c := nextInt(), nextInt()
	maxLen := 1<<31 - 1

	houses := make([]int, n)
	for i := range houses {
		houses[i] = nextInt()
	}

	sort.Ints(houses)

	l, r := 1, maxLen

	check := func(houses []int, mid, c int) bool {
		cnt := 1
		prev := houses[0]
		for i := 1; i < len(houses); i++ {
			if houses[i]-prev >= mid {
				cnt++
				prev = houses[i]
			}
		}
		return cnt >= c
	}

	for l <= r {
		mid := (l + r) / 2

		if check(houses, mid, c) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Fprintln(wr, r)
}
