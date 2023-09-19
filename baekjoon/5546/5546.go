package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	dp    [][][]int
	pasta []int
	N, K  int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N, K = nextInt(), nextInt()
	dp = make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([][]int, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = make([]int, 3)
		}
	}

	pasta = make([]int, N+1)

	for i := 0; i < K; i++ {
		a, b := nextInt(), nextInt()
		pasta[a] = b
	}

	if pasta[1] != 0 {
		dp[1][pasta[1]][1] = 1
	} else {
		dp[1][1][1] = 1
		dp[1][2][1] = 1
		dp[1][3][1] = 1
	}

	for i := 2; i < N+1; i++ {
		if pasta[i] != 0 {
			for j := 1; j < 4; j++ {
				if j == pasta[i] {
					continue
				}
				dp[i][pasta[i]][1] += dp[i-1][j][1] + dp[i-1][j][2]
				dp[i][pasta[i]][1] %= 10000
			}
			dp[i][pasta[i]][2] = dp[i-1][pasta[i]][1]
		} else {
			for j := 1; j < 4; j++ {
				dp[i][j][2] = dp[i-1][j][1]
				for k := 1; k < 4; k++ {
					if k == j {
						continue
					}
					dp[i][j][1] += dp[i-1][k][1] + dp[i-1][k][2]
					dp[i][j][1] %= 10000
				}
			}
		}
	}

	sum := dp[N][1][1] + dp[N][2][1] + dp[N][3][1] + dp[N][1][2] + dp[N][2][2] + dp[N][3][2]
	wr.WriteString(strconv.Itoa(sum % 10000))

}
