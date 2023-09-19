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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()

	stack := make([]int, 0, 10000)

	for i := 0; i < n; i++ {
		command := nextWord()
		switch command {
		case "push":
			stack = append(stack, nextInt())
		case "pop":
			if len(stack) == 0 {
				fmt.Fprintln(wr, -1)

			} else {
				fmt.Fprintln(wr, stack[0])
				stack = stack[1:]
			}
		case "size":
			fmt.Fprintln(wr, len(stack))
		case "empty":
			if len(stack) == 0 {
				fmt.Fprintln(wr, 1)
			} else {
				fmt.Fprintln(wr, 0)
			}
		case "front":
			if len(stack) == 0 {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, stack[0])
			}
		case "back":
			if len(stack) == 0 {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, stack[len(stack)-1])
			}
		}
	}
}
