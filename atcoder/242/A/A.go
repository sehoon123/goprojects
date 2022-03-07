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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a, b, c, x := nextInt(), nextInt(), nextInt(), nextInt()
	if x <= a {
		fmt.Println(1)
	} else if x > a && x <= b {
		fmt.Println(c / (b - a))
	} else {
		fmt.Println(0)
	}

}
