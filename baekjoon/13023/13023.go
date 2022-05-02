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
	adj     [][]int
	visited []bool
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
	adj = make([][]int, n)
	visited = make([]bool, n)

	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i, 1)
			visited[i] = false
		}
	}

	wr.WriteString("0\n")
}

func dfs(a int, cnt int) {
	if cnt == 5 {
		fmt.Println(1)
		os.Exit(0)
		return
	}

	for _, b := range adj[a] {
		if !visited[b] {
			visited[b] = true
			dfs(b, cnt+1)
			visited[b] = false
		}
	}
}
