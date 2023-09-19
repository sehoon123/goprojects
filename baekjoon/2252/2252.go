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
	adj     [][]int // smaller than index
	visited []bool
	answer  string
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	adj = make([][]int, n+1)
	visited = make([]bool, n+1)
	answer = ""

	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		adj[b] = append(adj[b], a)
	}

	//fmt.Fprintln(wr, adj)

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	fmt.Fprintln(wr, answer)

}

func dfs(x int) {
	visited[x] = true

	for _, v := range adj[x] {
		if !visited[v] {
			dfs(v)
		}
	}
	answer += strconv.Itoa(x) + " "

}
