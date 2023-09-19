package main

import "fmt"

func main() {
	var d1, d2 int
	fmt.Scan(&d1, &d2)
	dp := make([][]int, 2001)
	for i := 0; i <= 2000; i++ {
		dp[i] = make([]int, 2001)
	}

	result := 0

	for i := d1; i <= d2; i++ {
		for j := 0; j < i; j++ {
			k := gcd(i, j)
			a := i / k
			b := j / k
			if dp[a][b] == 0 {
				dp[a][b] = 1
				result++
			}
		}
	}

	fmt.Println(result)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
