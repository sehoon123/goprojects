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
	graph [][]int
	temp  [][]int
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
	n := nextInt()
	graph = make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = nextInt()
		}
	}

	//for _, v := range graph {
	//	fmt.Fprintln(wr, v)
	//}
	fmt.Fprintln(wr, solve(n))
}

func solve(n int) int {
	if n == 1 {
		return temp[0][0]
	}
	m := n / 2
	temp = make([][]int, m)
	for i := range temp {
		temp[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			test := make([]int, 4)
			test[0] = graph[2*i][2*j]
			test[1] = graph[2*i][2*j+1]
			test[2] = graph[2*i+1][2*j]
			test[3] = graph[2*i+1][2*j+1]
			sort.Ints(test)
			temp[i][j] = test[2]
		}
	}
	graph = temp
	return solve(n / 2)
}
