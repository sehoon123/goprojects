package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	cache []int
	pan   []int
	n, m  int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	cache = make([]int, 10001)
	pan = make([]int, m)
	for i := 0; i < m; i++ {
		pan[i] = nextInt()
		cache[pan[i]] = 1
	}

	for i := 1; i <= n; i++ {
		cache[i] = min(cache[i-])
	}

}
