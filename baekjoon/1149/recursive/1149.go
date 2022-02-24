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
	n     int
	color [][]int
	cache [][]int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n = nextInt()
	cache = make([][]int, n+1)
	color = make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		cache[i] = make([]int, n+1)
		r, g, b := nextInt(), nextInt(), nextInt()
		color[i] = []int{r, g, b}
	}

	a := solve(n, 0)
	b := solve(n, 1)
	c := solve(n, 2)

	if a <= b && a <= c {
		fmt.Println(a)
	} else if b <= a && b <= c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}
}

func solve(num, col int) int {
	if num == 1 {
		return color[num][col]
	}
	if cache[num][col] != 0 {
		return cache[num][col]
	}
	if col == 0 {
		cache[num][col] = min(solve(num-1, 1)+color[num][0], solve(num-1, 2)+color[num][0])
	} else if col == 1 {
		cache[num][col] = min(solve(num-1, 0)+color[num][1], solve(num-1, 2)+color[num][1])
	} else {
		cache[num][col] = min(solve(num-1, 1)+color[num][2], solve(num-1, 0)+color[num][2])
	}
	return cache[num][col]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
