package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	wr   = bufio.NewWriter(os.Stdout)
	ltoh [][]int
	htol [][]int
	n, m int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m = nextInt(), nextInt()
	ltoh = make([][]int, n+1)
	htol = make([][]int, n+1)
	for i := 0; i < m; i++ {
		h, l := nextInt(), nextInt()
		ltoh[l] = append(ltoh[l], h)
		htol[h] = append(htol[h], l)
	}

	count := 0
	for i := 1; i <= n; i++ {
		if bfs(i, ltoh) {
			count++
		}

		if bfs(i, htol) {
			count++
		}
	}
	fmt.Fprintln(wr, count)
}

func bfs(x int, arr [][]int) bool {
	visited := make([]int, n+1)
	visited[x] = 1
	temp := make([]int, 0, n)
	q := make([]int, 0, n)
	q = append(q, x)

	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		for _, v := range arr[now] {
			if visited[v] == 0 {
				q = append(q, v)
				visited[v] = 1
				temp = append(temp, v)
			}
		}
	}

	return len(temp) >= (n+1)/2
}
