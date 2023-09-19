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
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	T := nextInt()
	for i := 0; i < T; i++ {
		sum := 0
		min := 111
		for j := 0; j < 7; j++ {
			num := nextInt()
			if num%2 == 0 {
				sum += num
				if num < min {
					min = num
				}
			}
		}
		fmt.Fprintln(wr, sum, min)
	}
}
