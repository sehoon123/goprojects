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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()
	k := nextInt()

	graph := make([][]byte, n)
	dp := make([][][3]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][3]int, m+1)
	}
	//fmt.Fprintln(wr, dp)

	for i := 0; i < n; i++ {
		temp := nextWord()
		graph[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			graph[i][j] = temp[j]
		}
	}

	for i := 0; i < n; i++ {
		J, O, I := 0, 0, 0
		for j := 0; j < m; j++ {
			if graph[i][j] == 'J' {
				J++
			} else if graph[i][j] == 'O' {
				O++
			} else if graph[i][j] == 'I' {
				I++
			}
			dp[i+1][j+1][0] = J
			dp[i+1][j+1][1] = O
			dp[i+1][j+1][2] = I
		}
	}

	for i := 0; i < k; i++ {
		a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
		J, O, I := 0, 0, 0
		for j := a; j < c+1; j++ {
			J += dp[j][d][0] - dp[j][b-1][0]
			O += dp[j][d][1] - dp[j][b-1][1]
			I += dp[j][d][2] - dp[j][b-1][2]
		}
		fmt.Fprintln(wr, J, O, I)
	}

}
