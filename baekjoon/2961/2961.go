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

	arr := make([][2]int, n)

	for i := 0; i < n; i++ {
		arr[i][0] = nextInt()
		arr[i][1] = nextInt()
	}

	bit := 0
	min := abs(arr[0][0] - arr[0][1])

	for bit = 1; bit < 1<<uint(n); bit++ {
		sour := 1
		bitt := 0
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) != 0 {
				sour *= arr[i][0]
				bitt += arr[i][1]
			}
		}
		if abs(sour-bitt) < min {
			min = abs(sour - bitt)
		}
	}

	fmt.Fprintln(wr, min)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
