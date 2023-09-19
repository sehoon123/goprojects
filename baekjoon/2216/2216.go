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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	A, B, C := nextInt()*-1, nextInt()*-1, nextInt()*-1

	X, Y := nextWord(), nextWord()

	dp := make([][]int, len(X)+1)
	for i := 0; i < len(X)+1; i++ {
		dp[i] = make([]int, len(Y)+1)
	}

	for i := 1; i <= len(X); i++ {
		dp[i][0] = dp[i-1][0] + B
	}
	for i := 1; i <= len(Y); i++ {
		dp[0][i] = dp[0][i-1] + B
	}

	for i := 1; i <= len(X); i++ {
		for j := 1; j <= len(Y); j++ {
			temp := 0
			if X[i-1] == Y[j-1] {
				temp = dp[i-1][j-1] + A
			} else {
				temp = dp[i-1][j-1] + C
			}

			dp[i][j] = min(dp[i-1][j]+B, dp[i][j-1]+B, temp)
		}
	}

	fmt.Fprintln(wr, dp)
	fmt.Fprintln(wr, dp[len(X)][len(Y)]*-1)

}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
