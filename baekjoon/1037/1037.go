package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	k := nextInt()
	divisor := make([]int, k)
	for i := 0; i < k; i++ {
		divisor[i] = nextInt()
	}

	sort.Ints(divisor)
	fmt.Fprintln(wr, divisor[0]*divisor[k-1])

}
