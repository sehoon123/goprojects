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

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()
	graph := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			graph[i][j] = graph[i][j-1] + nextInt()
		}
	}

	req := make([][]int, m)
	for i := 0; i < m; i++ {
		req[i] = make([]int, 4)
		req[i][0], req[i][1], req[i][2], req[i][3] = nextInt(), nextInt(), nextInt(), nextInt()
	}

	for i := 0; i < m; i++ {
		temp := 0
		for x := req[i][0]; x <= req[i][2]; x++ {
			temp += graph[x][req[i][3]] - graph[x][req[i][1]-1]
		}
		fmt.Fprintln(wr, temp)
	}

}
