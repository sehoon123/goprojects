package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	word := nextString()

	arr := make([]int, 2)
	now := 0

	for _, v := range word {
		if v == '@' {
			arr[now]++
		}
		if v == '0' {
			now = 1
		}
	}

	fmt.Fprintln(wr, arr[0], arr[1])

}
