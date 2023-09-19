package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	defer wr.Flush()

	for {
		sc.Scan()
		line := sc.Text()
		line = strings.ToUpper(line)

		if line == "#" {
			break
		}

		ans := 0
		count := make(map[rune]int)

		for _, r := range line {
			count[r]++
			if r >= 'A' && r <= 'Z' && count[r] == 1 {
				ans++
			}
		}

		fmt.Fprintln(wr, ans)
	}

}
