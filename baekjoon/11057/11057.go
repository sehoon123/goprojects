package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 0; j < 10; j++ {
			sum := 0
			for k := j; k < 10; k++ {
				sum += dp[k]
			}
			dp[j] = sum % 10007
		}
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += dp[i]
	}

	fmt.Println(sum % 10007)
}
