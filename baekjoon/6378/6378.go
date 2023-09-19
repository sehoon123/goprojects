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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	for {
		n := nextWord()
		if n == "0" {
			break
		}
		fmt.Fprintln(wr, digitalRoot(n))

	}
}

func digitalRoot(n string) int {
	if len(n) == 1 {
		return int(n[0] - '0')
	}
	temp := 0
	for _, v := range n {
		temp += int(v - '0')
	}
	return digitalRoot(fmt.Sprint(temp))

}
