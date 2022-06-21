package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	cycle []int
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	v, e := nextInt(), nextInt()

	node := make([][]int, e)
	cycle = make([]int, v+1)

	for i := 1; i <= v; i++ {
		cycle[i] = i
	}

	for i := 0; i < e; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		node[i] = []int{a, b, c}
	}

	sort.Slice(node, func(i, j int) bool {
		return node[i][2] < node[j][2]
	})

	//fmt.Fprintln(wr, node)

	sum := 0
	for i := 0; i < e; i++ {
		if !findParent(node[i][0], node[i][1]) {
			sum += node[i][2]
			unionParent(node[i][0], node[i][1])
		}
	}

	fmt.Fprintln(wr, sum)
}

func getParent(x int) int {
	if cycle[x] == x {
		return x
	}
	cycle[x] = getParent(cycle[x])
	return cycle[x]
}

func unionParent(x, y int) {
	x, y = getParent(x), getParent(y)
	if x < y {
		cycle[y] = x
	} else {
		cycle[x] = y
	}
}

func findParent(x, y int) bool {
	x, y = getParent(x), getParent(y)
	if x == y {
		return true
	}
	return false
}
