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
		n := nextInt()
		candidates := make([]int, n+1)
		sum := 0
		max := 0
		min := int(1e9)
		idx := 0
		man := 0
		for j := 1; j <= n; j++ {
			candidates[j] = nextInt()
			if candidates[j] == max {
				man += 1
			}
			if candidates[j] > max {
				max = candidates[j]
				idx = j
				man = 0
			} else if candidates[j] < min {
				min = candidates[j]
			}
			sum += candidates[j]
		}

		if max-min == 0 || man > 0 {
			fmt.Fprintln(wr, "no winner")
		} else if max > sum/2 {
			fmt.Fprintln(wr, "majority winner", idx)
		} else {
			fmt.Fprintln(wr, "minority winner", idx)
		}
	}
}
