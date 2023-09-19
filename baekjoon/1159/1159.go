package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextWord() uint8 {
	sc.Scan()
	return sc.Text()[0]
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	var n int
	fmt.Scan(&n)

	player := make(map[uint8]int)
	for i := 0; i < n; i++ {
		player[nextWord()]++
	}

	result := make([]string, 0)

	for k, v := range player {
		if v >= 5 {
			result = append(result, string(k))
		}
	}

	sort.Strings(result)

	if len(result) == 0 {
		fmt.Fprintln(wr, "PREDAJA")
	} else {
		for _, v := range result {
			fmt.Fprintf(wr, "%v", v)
		}
	}

}
