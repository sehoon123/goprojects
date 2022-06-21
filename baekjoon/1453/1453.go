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
	ans := 0
	var list map[int]int
	list = make(map[int]int)

	for i := 0; i < n; i++ {
		num := nextInt()
		if list[num] == 1 {
			ans++
		} else {
			list[num] = 1
		}
	}

	fmt.Fprintln(wr, ans)
}
