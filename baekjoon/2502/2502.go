package main

import "fmt"

func main() {
	var d, k int
	fmt.Scan(&d, &k)
	dp := make([]int, d+1)
	for i := 1; i <= k; i++ {
		dp[1] = i
		for j := i; j <= k; j++ {
			dp[2] = j
			for l := 3; l <= d; l++ {
				dp[l] = dp[l-1] + dp[l-2]
			}
			if dp[d] == k {
				fmt.Printf("%d\n%d", i, j)
				return
			}
			if dp[d] > k {
				break
			}

		}
	}
}
