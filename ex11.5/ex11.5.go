package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//별찍기 연습, 마름모, 가운데정렬 별 찍기
func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	for i := 1; i <= n; i++ {
		fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", n-i), strings.Repeat("*", i*2-1))
	}
	for i := n-1; i >= 1; i-- {
		fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", n-i), strings.Repeat("*", i*2-1))
	}
}
