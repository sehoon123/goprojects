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
	for i := 0; i < m; i++ {
		k := 0
		if line[i] == 'O' {
			continue
		} else {
			for line[i+1] == 'O' && line[i+2] == 'I' {
				k++
				if k == n {
					k--
					count++
				}
				i += 2
			}
			k = 0
		}
	}

	fmt.Fprintln(w, count)

}
