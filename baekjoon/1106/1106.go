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
	c, n  int
	price [][]int
	cache []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	cache = make([]int, 1001)
	for i := range cache {
		cache[i] = int(1e9)
	}
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	c, n = nextInt(), nextInt()
	price = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		price = append(price, []int{a, b})
	}
	fmt.Fprintln(wr, sol(c))
}

func sol(cnt int) int {
	if cnt < 1 {
		return 0
	}
	if cache[cnt] != int(1e9) {
		return cache[cnt]
	}
	for _, v := range price {
		cache[cnt] = min(cache[cnt], sol(cnt-v[1])+v[0])
	}
	return cache[cnt]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
