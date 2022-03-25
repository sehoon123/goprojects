package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m, B := nextInt(), nextInt(), nextInt()
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			graph[i][j] = nextInt()
		}
	}
	max_height := 0
	min_time := 0

loop:
	for height := 0; height <= 256; height++ {
		var b int
		b = B

		temp := make([][]int, n)
		for i := 0; i < n; i++ {
			temp[i] = make([]int, m)
			for j := 0; j < m; j++ {
				temp[i][j] = graph[i][j]
			}
		}
		var time int

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if temp[i][j] > height {
					b += temp[i][j] - height
					time += 2 * (temp[i][j] - height)
					temp[i][j] = height
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if temp[i][j] < height {
					b -= height - temp[i][j]
					if b < 0 {
						break loop
					}
					time += 1 * (height - temp[i][j])
					temp[i][j] = height
				}
			}
		}
		if time <= min_time || min_time == 0 {
			min_time = time
			max_height = height
		}
	}
	fmt.Fprintln(wr, min_time, max_height)

}
