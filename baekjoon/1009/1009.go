package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r, w := bufio.NewReader(bufio.NewReader(os.Stdin)), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		result := math.Pow(float64(a), float64(b%4+4))
		fmt.Fprintln(w, result)
	}
}
