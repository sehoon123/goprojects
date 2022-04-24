package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	n := readInt()
	a := make([]int64, n+1)
	a[0] = math.MinInt64
	b := a[1:]
	readInt64s(&b)

	d := make([]int, n+1)
	x := make([]int64, 1, n+1)
	x[0] = math.MinInt64

	for i := 1; i <= n; i++ {
		pos := sort.Search(len(x), func(j int) bool { return x[j] >= a[i] })
		d[i] = pos
		if pos >= len(x) {
			x = append(x, a[i])
		} else {
			if x[pos] > a[i] {
				x[pos] = a[i]
			}
		}
	}

	res := x[1:]
	fmt.Fprintln(writer, res)
}

func readInt() int {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	res, _ := strconv.Atoi(text)
	return res
}

func readInt64s(dst *[]int64) {
	text, _ := reader.ReadString('\n')
	fld := strings.Fields(text)
	for i := range *dst {
		(*dst)[i], _ = strconv.ParseInt(fld[i], 10, 64)
	}
}
