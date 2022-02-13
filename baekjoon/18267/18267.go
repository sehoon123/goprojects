package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)
var str []byte
var adj [][]int
var comp []int
var num int
var s string

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextChar() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 1024*1024)

	wr = bufio.NewWriterSize(os.Stdout, 2024*1024)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m := nextInt(), nextInt()
	adj = make([][]int, n+1)
	comp = make([]int, n+1)

	sc.Scan()
	s = sc.Text()

	for i := 0; i < n-1; i++ {
		start, end := nextInt(), nextInt()
		adj[start] = append(adj[start], end)
		adj[end] = append(adj[end], start)
	}

	num = 0
	for i := 1; i < n+1; i++ {
		if comp[i] == 0 {
			num += 1
			dfs(i)
		}
	}

	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextChar()
		if string(s[a-1]) == c || comp[a] != comp[b] {
			fmt.Fprintf(wr, "%d", 1)
		} else {
			fmt.Fprintf(wr, "%d", 0)
		}
	}
}

func dfs(x int) {
	if comp[x] != 0 {
		return
	}
	comp[x] = num
	for _, v := range adj[x] {
		if s[v-1] == s[x-1] {
			dfs(v)
		}
	}
}
