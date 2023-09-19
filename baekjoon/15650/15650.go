package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc      = bufio.NewScanner(os.Stdin)
	wr      = bufio.NewWriter(os.Stdout)
	N, M    int
	visited []bool
	arr     []int
	cnt     int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N, M = nextInt(), nextInt()
	visited = make([]bool, N+1)
	arr = make([]int, 0, M)

	dfs(N, M)
}

func dfs(n, m int) {
	if len(arr) == M {
		for _, v := range arr {
			fmt.Fprintf(wr, "%d ", v)
		}
		fmt.Fprintln(wr)
		return
	}

	for i := 1; i < N+1; i++ {
		if !visited[i] {
			arr = append(arr, i)
			dfs(n, m)
			visited[i] = true
			arr = arr[:len(arr)-1]
			visited[i] = false
		}
	}
}
