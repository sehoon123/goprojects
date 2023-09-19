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

func nextInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	max := 0
	x, y := 1, 1

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			num := nextInt()
			if num > max {
				max = num
				x, y = i, j
			}
		}
	}

	fmt.Fprintln(wr, max)
	fmt.Fprintln(wr, x, y)
}
