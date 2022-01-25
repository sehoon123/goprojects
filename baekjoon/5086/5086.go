package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for {
		var n, m int
		fmt.Fscan(r, &n, &m)
		if n == 0 && m == 0 {
			break
		}
		if m%n == 0 {
			fmt.Fprintln(w, "factor")
		} else if n%m == 0 {
			fmt.Fprintln(w, "multiple")
		} else {
			fmt.Fprintln(w, "neither")
		}

	}
}
