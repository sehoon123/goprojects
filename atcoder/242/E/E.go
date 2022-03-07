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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	T := nextInt()
	for t := 0; t < T; t++ {
		n := nextInt()
		s := nextString()
		char := make([]rune, n)
		for i := 0; i < n; i++ {
			char[i] = rune(s[i] - 'A')
		}
		fmt.Println(char)
	}
}
