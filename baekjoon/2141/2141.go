package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc      = bufio.NewScanner(os.Stdin)
	wr      = bufio.NewWriter(os.Stdout)
	village [][2]int
	n       int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	n = nextInt()
	// village[i][0] : 마을위치 , village[i][1] : 사람 수
	village = make([][2]int, n)

	for i := 0; i < n; i++ {
		village[i][0], village[i][1] = nextInt(), nextInt()
	}

	sort.Slice(village, func(i, j int) bool {
		return village[i][0] < village[j][0]
	})

	left := village[0][0]
	right := village[n-1][0]
	res := int(1e12)
	P := int(1e12)

	for left <= right {
		mid := (left + right) / 2
		dis := f(mid)
		s := dis[0] + dis[1]
		if s < res {
			res = s
			P = mid
		} else if s == res {
			P = min(mid, P)
		}
		if dis[0] > dis[1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Fprintln(wr, P)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(mid int) [2]int {
	left, right := 0, 0
	for i := 0; i < n; i++ {
		x, y := village[i][0], village[i][1]
		if x < mid {
			left += y
		} else if x > mid {
			right += y
		}
	}
	return [2]int{left, right}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
