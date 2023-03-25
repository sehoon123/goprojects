package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dp []int
)

func nextInt() int {
	sc.Scan()
	r := 0
	for _, c := range sc.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	dp = make([]int, n+1)
	dp[1] = 0
	dp[2] = 3

	for i := 4; i <= n; i += 2 {
		dp[i] = dp[i-2] * dp[2]
		for j := 4; j <= i; j += 2 {
			dp[i] += dp[i-j] * 2 // 2는 짝수일때 생기는 특이한 모양에 대응하는 숫자임
		}
		dp[i] += 2
	}

	fmt.Fprintln(wr, dp[n])
}
