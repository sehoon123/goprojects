package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var line1, line2 string
	fmt.Fscanln(r, &line1)
	fmt.Fscanln(r, &line2)

	fmt.Fprintln(w, lcs(line1, line2))
}

//longest common subsequence
func lcs(a, b string) int {
	var dp [][]int
	for i := 0; i <= len(a); i++ {
		dp = append(dp, make([]int, len(b)+1))
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(a)][len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
