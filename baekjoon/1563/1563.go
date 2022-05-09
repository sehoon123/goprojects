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
	dp [][2][3]int
)

const (
	MOD = 1000000
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	// dp[date][late][absent continue]
	dp = make([][2][3]int, n+1)

	dp[1][0][0], dp[1][0][1], dp[1][1][0] = 1, 1, 1

	for i := 2; i <= n; i++ {
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % MOD
		dp[i][0][1] = dp[i-1][0][0] % MOD
		dp[i][0][2] = dp[i-1][0][1] % MOD
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2] + dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2]) % MOD
		dp[i][1][1] = (dp[i-1][1][0]) % MOD
		dp[i][1][2] = (dp[i-1][1][1]) % MOD
	}

	fmt.Fprintln(wr, (dp[n][0][0]+dp[n][0][1]+dp[n][0][2]+dp[n][1][0]+dp[n][1][1]+dp[n][1][2])%MOD)
}
