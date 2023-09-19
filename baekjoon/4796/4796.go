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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	i := 1
	for {
		a, b, c := nextInt(), nextInt(), nextInt()
		if a == 0 && b == 0 && c == 0 {
			break
		}

		result := 0
		if c%b > a {
			result = int(c/b)*a + a
		} else {
			result = int(c/b)*a + c%b
		}
		fmt.Fprintf(wr, "%s %d: %d\n", "Case", i, result)
		i++
	}
}
