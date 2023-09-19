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
	arr := [10]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	cheat := make([]int, 0)

	for i := 1; i <= n; i++ {
		temp := [10]int{}
		for j := 0; j < 10; j++ {
			temp[j] = nextInt()
		}

		if temp == arr {
			cheat = append(cheat, i)
		}
	}

	for _, v := range cheat {
		fmt.Fprintln(wr, v)
	}
}
