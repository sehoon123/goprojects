package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	const mod = 1234567

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		dp := make([]int, 10)
		for j := 0; j < 10; j++ {
			dp[j] = 1
		}

		for j := 1; j < n; j++ {
			temp := make([]int, 10)
			copy(temp, dp)
			temp[1] = (dp[2] + dp[4]) % mod
			temp[2] = (dp[1] + dp[3] + dp[5]) % mod
			temp[3] = (dp[2] + dp[6]) % mod
			temp[4] = (dp[1] + dp[5] + dp[7]) % mod
			temp[5] = (dp[2] + dp[4] + dp[6] + dp[8]) % mod
			temp[6] = (dp[3] + dp[5] + dp[9]) % mod
			temp[7] = (dp[4] + dp[8] + dp[0]) % mod
			temp[8] = (dp[5] + dp[7] + dp[9]) % mod
			temp[9] = (dp[6] + dp[8]) % mod
			temp[0] = (dp[7]) % mod
			copy(dp, temp)
		}

		fmt.Println(sum(dp) % mod)

	}
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
