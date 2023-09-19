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
	arr = make([]int, 0, M)
	visited = make([]bool, N+1)

	dfs(N, M)
}

func dfs(N, M int) {
	if len(arr) == M {
		for _, v := range arr {
			fmt.Fprintf(wr, "%d ", v)
		}
		fmt.Fprintln(wr)
		return
	}

	for i := 1; i <= N; i++ {
		if !visited[i] {
			arr = append(arr, i)
			dfs(N, M)
			visited[i] = true
			arr = arr[:len(arr)-1]
			for j := i + 1; j <= N; j++ {
				visited[j] = false
			}
		}
	}
}
