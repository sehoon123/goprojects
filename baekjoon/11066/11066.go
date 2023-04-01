package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

const INF = math.MaxInt64

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	T := nextInt()

	for t := 0; t < T; t++ {
		K := nextInt()

		dp := make([][]int, K+1)
		sum := make([]int, K+1)
		arr := make([]int, K+1)

		for i := 1; i <= K; i++ {
			arr[i] = nextInt()
			dp[i] = make([]int, K+1)
			sum[i] = sum[i-1] + arr[i]
			dp[i][i] = 0
			if i < K {
				dp[i][i+1] = arr[i] + arr[i+1]
			}
		}

		for i := range dp {
			fmt.Fprintln(wr, dp[i])
		}

		for d := 2; d <= K; d++ {
			for i := 1; i <= K-d+1; i++ {
				j := i + d - 1
				dp[i][j] = INF

				for k := i; k <= j-1; k++ {
					cost := dp[i][k] + dp[k+1][j] + sum[j] - sum[i-1]
					if dp[i][j] > cost {
						dp[i][j] = cost
					}
				}

				for i := range dp {
					fmt.Fprintln(wr, dp[i])
				}
			}
		}

		fmt.Fprintln(wr, dp[1][K])
	}
}
