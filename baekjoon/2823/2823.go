package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	r, c := nextInt(), nextInt()
	graph := make([]string, r)
	for i := 0; i < r; i++ {
		sc.Scan()
		graph[i] = sc.Text()
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	count := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if graph[i][j] == '.' {
				for k := 0; k < 4; k++ {
					nx, ny := i+dx[k], j+dy[k]
					if nx < 0 || nx >= r || ny < 0 || ny >= c {
						count++
						continue
					}
					if graph[nx][ny] == 'X' {
						count++
					}
				}
				if count >= 3 {
					fmt.Fprintln(wr, 1)
					return
				}
				count = 0
			}
		}
	}

	fmt.Fprintln(wr, 0)
}
