package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() uint64 {
	sc.Scan()
	i, _ := strconv.ParseUint(sc.Text(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	a, b := nextInt(), nextInt()
	dp := make([]uint64, 100)
	dp[0] = 1
	for i := 1; i < 100; i++ {
		dp[i] = dp[i-1]*2 + (1 << uint64(i))
	}

	fmt.Fprintln(wr, dp[0:10])
	fmt.Fprintln(wr, a, b)

}

func solve(a uint64) uint64 {
	var tmp uint64 = a
	for tmp != 0 {

	}
}
