package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()

	for i := 0; i < n; i++ {
		visited := make([]bool, 100)
		word := nextWord()
		total := 1950 + 65
		for _, v := range word {
			if !visited[v-'A'] {
				total -= int(v)
				visited[v-'A'] = true
			}
		}
		wr.WriteString(strconv.Itoa(total) + "\n")
	}
}
