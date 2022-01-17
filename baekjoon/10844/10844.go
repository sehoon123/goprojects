package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	d := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = make([]int, 10)
	}

	for i := 1; i < 10; i++ {
		d[1][i] = 1
	}

	for i := 2; i < n+1; i++ {
		d[i][0] = d[i-1][1]
		d[i][9] = d[i-1][8]

		for j := 1; j < 9; j++ {
			d[i][j] = int(d[i-1][j-1]+d[i-1][j+1]) % 1e9
		}
	}
	sum := 0
	for _, v := range d[n] {
		sum += v
	}
	fmt.Fprintln(w, sum%1e9)
}
