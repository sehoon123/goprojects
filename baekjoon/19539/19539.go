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
	sum := 0
	temp := 0

	for i := 0; i < n; i++ {
		num := nextInt()
		sum += num
		temp += num % 2
	}

	if sum%3 != 0 {
		fmt.Fprintln(wr, "NO")
	} else if sum/3 >= temp {
		fmt.Fprintln(wr, "YES")
	} else {
		fmt.Fprintln(wr, "NO")
	}
}
