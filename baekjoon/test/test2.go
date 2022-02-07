package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	graph := make([]string, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(r, &graph[i])
	}

	fmt.Fprintln(w, graph[1][1] == '1')

}
