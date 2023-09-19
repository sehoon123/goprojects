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
	n := nextInt()
	score := make([]int, 100000000)
	for i := 0; i < n*n; i++ {
		s, k := nextInt(), nextInt()
		for j := 0; j < k; j++ {
			t := nextInt()
			if score[t] < s {
				score[t] = s
			}
		}
	}
	ans := 0
	for i := 0; i < 100000000; i++ {
		ans += score[i]
	}
	fmt.Fprintln(wr, ans)

}
