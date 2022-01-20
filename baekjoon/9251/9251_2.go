package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s1, s2 string
	fmt.Fscanln(r, &s1)
	fmt.Fscanln(r, &s2)

	//fmt.Fprintln(w, len(s1))
	d := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		d[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				d[i][j] = d[i-1][j-1] + 1
			} else {
				d[i][j] = max(d[i-1][j], d[i][j-1])
			}
		}
	}

	for _, v := range d {
		fmt.Fprintln(w, v)
	}

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
