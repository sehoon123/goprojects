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
	graph := make([]string, n)
	rank := make([]int, 10)
	score := make([]int, 51)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &graph[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 1; j < m-1; j++ {
			if graph[i][j] != '.' {
				rank[graph[i][j]-'0'] = count
				break
			} else {
				count++
			}
		}
		count = 0
	}

	temp := make([]int, 10)
	copy(temp, rank)

	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	a := 1
	for _, v := range temp {
		if score[v] == 0 {
			score[v] = a
			a++
		}
	}

	for i := 1; i <= 9; i++ {
		fmt.Fprintln(w, score[rank[i]])
	}

}
