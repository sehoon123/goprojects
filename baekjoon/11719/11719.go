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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	//sc.Split(bufio.ScanWords)
	defer wr.Flush()

	for sc.Scan() {
		fmt.Fprintln(wr, sc.Text())
	}
}
