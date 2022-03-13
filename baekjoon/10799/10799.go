package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	paren []byte
)

func nextByte() []byte {
	sc.Scan()
	return sc.Bytes()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	paren = nextByte()
	//stack := make([]int, 0, len(paren)+1)
	stack := []int{}
	result := 0

	for i := range paren {
		if paren[i] == '(' {
			stack = append(stack, 1)
		} else {
			stack = stack[:len(stack)-1]
			if paren[i-1] == '(' {
				result += len(stack)
			} else {
				result++
			}
		}
	}

	fmt.Fprintln(wr, result)
}
