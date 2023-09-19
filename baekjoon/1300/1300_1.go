package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	start, end := 1, k
	for start <= end {
		mid := (start + end) / 2
		count := 0
		for i := 1; i < n+1; i++ {
			count += min(mid/i, n)
		}
		if count >= k {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	fmt.Println(start)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
