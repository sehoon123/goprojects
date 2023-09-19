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

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	m := nextInt()
	bit := 1 << 20

	for i := 0; i < m; i++ {
		//fmt.Printf("%b", bit)
		command := nextWord()
		switch command {
		case "add":
			x := nextInt()
			bit |= 1 << (x - 1)
		case "check":
			x := nextInt()
			if bit&(1<<(x-1)) != 0 {
				fmt.Fprintln(wr, 1)
			} else {
				fmt.Fprintln(wr, 0)
			}
		case "remove":
			x := nextInt()
			bit &= ^(1 << (x - 1))
		case "toggle":
			x := nextInt()
			bit ^= (1 << (x - 1))
		case "all":
			bit = (1 << 21) - 1
		case "empty":
			bit = 0
		}

	}

}
