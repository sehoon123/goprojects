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
	cahce [][]int
	n     int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	for {
		n = nextInt()
		if n == 0 {
			break
		}
		cahce = make([][]int, n+1)
		for i := 0; i < n+1; i++ {
			cahce[i] = make([]int, n+1)
		}
		fmt.Fprintln(wr, solve(0, 0))
	}
}

func solve(O, X int) int {
	if O == n {
		return 1
	}
	if X > O {
		return 0
	}
	if cahce[O][X] != 0 {
		return cahce[O][X]
	}
	cahce[O][X] = solve(O, X+1) + solve(O+1, X)

	return cahce[O][X]
}
