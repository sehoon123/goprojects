package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n, k := nextInt(scanner), nextInt(scanner)

	coins := make(map[int]bool)
	for i := 0; i < n; i++ {
		coin := nextInt(scanner)
		coins[coin] = true
	}

	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = 1e9
	}

	ans := solve(dp, coins, k)

	if ans == 1e9 {
		ans = -1
	}

	fmt.Println(ans)
}

func solve(dp []int, coins map[int]bool, n int) int {
	if n == 0 {
		return 0
	}
	if dp[n] != 1e9 {
		return dp[n]
	}
	for coin := range coins {
		if n-coin >= 0 {
			dp[n] = min(dp[n], solve(dp, coins, n-coin)+1)
		}
	}
	return dp[n]
}

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num := 0
	neg := false
	for _, c := range scanner.Bytes() {
		if c == '-' {
			neg = true
		} else {
			num = num*10 + int(c-'0')
		}
	}
	if neg {
		num = -num
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
