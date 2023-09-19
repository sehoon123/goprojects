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
	flag := true
	for i := 1; i < n+1; i++ {
		temp := nextInt()
		if i%2 == 1 {
			if temp%2 == 0 {
				flag = false
				break
			}
		} else if i%2 == 0 {
			if temp%2 == 1 {
				flag = false
				break
			}
		}
	}

	if flag {
		fmt.Fprintln(wr, "YES")
	} else {
		fmt.Fprintln(wr, "NO")
	}

}
