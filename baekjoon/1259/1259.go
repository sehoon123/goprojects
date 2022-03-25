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

func nextBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	for {
		str := nextBytes()
		if str[0] == '0' {
			break
		}

		for {
			if len(str) == 0 || len(str) == 1 {
				fmt.Fprintln(wr, "yes")
				break
			}

			if str[0] == str[len(str)-1] {
				str = str[1 : len(str)-1]
			} else {
				fmt.Fprintln(wr, "no")
				break
			}
		}
	}
}
