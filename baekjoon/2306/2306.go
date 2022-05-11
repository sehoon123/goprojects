package main

import "fmt"

func main() {
	var gene string
	fmt.Scan(&gene)
	n := len(gene)

	var m map[rune]rune
	m = make(map[rune]rune)
	m['a'] = 't'
	m['g'] = 'c'

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 1; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if (gene[i] == 'a' || gene[i] == 'g') && (rune(gene[j]) == m[rune(gene[i])]) {
				dp[i][j] = dp[i+1][j-1] + 2
			}
			for k := i; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k+1][j])
			}
		}
	}
	//for _, v := range dp {
	//	fmt.Println(v)
	//}
	fmt.Println(dp[0][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
//  PYTHON CODE
//
//buff = input()
//
//memo = [[0 for i in range(501)] for j in range(501)]
//
//for j in range(1,len(buff)):
//for i in range(j-1,-1,-1):
//if buff[i]+buff[j] in ('at','gc'):
//memo[i][j] = memo[i+1][j-1]+2
//for t in range(i+1,j):
//memo[i][j] = max(memo[i][j],memo[i][t]+memo[t][j])
//
//print( memo[0][len(buff)-1] )
