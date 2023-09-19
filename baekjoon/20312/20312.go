package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	table [][]int
)

const mod int = 1e9 + 7

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	arr := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		arr[i] = nextInt()
	}

	result := 0
	for i := 0; i < n-1; i++ {
		temp := 1
		for j := i; j < n-1; j++ {
			temp *= (arr[j]) % mod
			result += (temp) % mod
		}
	}
	fmt.Fprintln(wr, result%mod)

}
