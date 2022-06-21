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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a, b := nextWord(), nextWord()

	rev_a := ""
	rev_b := ""

	for _, v := range a {
		rev_a = string(v) + rev_a
	}

	for _, v := range b {
		rev_b = string(v) + rev_b
	}

	reva, _ := strconv.Atoi(rev_a)
	revb, _ := strconv.Atoi(rev_b)

	ans := reva + revb
	temp := strconv.Itoa(ans)

	res := ""
	for _, v := range temp {
		res = string(v) + res
	}

	final, _ := strconv.Atoi(res)

	fmt.Fprintln(wr, final)

}
