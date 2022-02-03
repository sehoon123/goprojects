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
	var n, m int
	fmt.Fscan(r, &n, &m)
	s := make([]byte, m+4)
	for i := 0; i < m+2; i++ {
		fmt.Fscanf(r, "%c", &s[i])
	}
	fmt.Fprintln(w, s)

	//	count := 0
	//	ans := 0
	//	for i := 2; i < m+4; {
	//		if s[i-2] == 'I' && s[i-1] == 'O' && s[i] == 'I' {
	//			count++
	//			i += 2
	//			if count == n {
	//				ans++
	//				count--
	//			}
	//		} else {
	//			count = 0
	//			i++
	//		}
	//	}
	//
	//	fmt.Fprintln(w, ans)
}
