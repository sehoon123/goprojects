package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscan(r, &n, &m)
	var line string
	fmt.Fscan(r, &line)

	count := 0
	ans := 0
	i := 0

	for i < m-2 {
		if line[i:i+3] == "IOI" {
			i += 2
			count++
			if count == n {
				count--
				ans++
			}
		} else {
			i++
			count = 0
		}

	}

	fmt.Fprintln(w, ans)

}
