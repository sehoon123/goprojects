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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n, m, k := nextInt(), nextInt(), nextInt()
	bitmask := 1<<m - 1
	preference := make([]int, n)

	for i := 0; i < n; i++ {
		preference[i] = 1<<m - 1
		for j := m - 1; j >= 0; j-- {
			x := nextInt()
			if x < 5 {
				preference[i] ^= 1 << j
			}
		}
	}

	result := 0
	for i := 0; i < bitmask; i++ {
		count := 0
		if bitcount(i) != m-k {
			continue
		}
		for j := 0; j < n; j++ {
			if i&preference[j] != 0 {
				count++
			}
		}
		if count > result {
			result = count
		}
	}
	fmt.Fprintln(wr, result)
}

func bitcount(bit int) int {
	if bit == 0 {
		return 0
	}
	return bit%2 + bitcount(bit/2)
}
