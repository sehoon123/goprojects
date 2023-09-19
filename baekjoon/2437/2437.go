package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	sort.Ints(a)
	dp := make([]int, 10000000)
	sum := 0

	for i := 0; i < n; i++ {
		for j := 1; j < sum+1; j++ {
			if dp[j] == 1 {
				dp[j+a[i]] = 1
			} else {
				break
			}
		}
		dp[a[i]] = 1
		sum += a[i]
	}

	for i := 1; i < 1000000010; i++ {
		if dp[i] == 0 {
			fmt.Fprintln(wr, i)
			break
		}
	}
}
