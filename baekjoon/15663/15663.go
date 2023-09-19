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
	N, M    int
	arr     []int
	temp    []int
	visited []bool
)

func nextInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	N, M = nextInt(), nextInt()

	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = nextInt()
	}

	sort.Ints(arr)
	temp = make([]int, M)
	visited = make([]bool, N)

	dfs(0, 0)

}

func dfs(idx, cnt int) {
	if idx == M {
		for i := 0; i < M; i++ {
			fmt.Fprintf(wr, "%d ", temp[i])
		}
		fmt.Fprintln(wr)
		return
	}

	before := -1
	for i := cnt; i < N; i++ {
		if !visited[i] && arr[i] != before {
			visited[i] = true
			temp[idx] = arr[i]
			before = arr[i]
			dfs(idx+1, cnt+1)
			visited[i] = false
		}
	}
}
