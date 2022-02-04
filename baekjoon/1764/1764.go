package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscan(r, &n, &m)
	bad := make(map[string]int, n)
	badass := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(r, &s)
		bad[s] = 0
	}
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(r, &s)
		if _, ok := bad[s]; ok {
			badass = append(badass, s)
		}
	}
	sort.Strings(badass)
	fmt.Fprintln(w, len(badass))
	for _, v := range badass {
		fmt.Fprintln(w, v)
	}
}
