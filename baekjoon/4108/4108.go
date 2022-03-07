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

func nextChar() []byte {
	sc.Scan()
	return sc.Bytes()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for {
		r, c := nextInt(), nextInt()
		result := make([][]string, r)
		for i := 0; i < r; i++ {
			result[i] = make([]string, c)
		}
		if r == 0 && c == 0 {
			break
		}
		graph := make([][]byte, r)
		for i := 0; i < r; i++ {
			graph[i] = nextChar()
		}
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				count := 0
				if graph[i][j] == '.' {
					for k := 0; k < 8; k++ {
						nx, ny := i+dx[k], j+dy[k]
						if nx < 0 || nx >= r || ny < 0 || ny >= c {
							continue
						}
						if graph[nx][ny] == '*' {
							count++
						}
					}
					result[i][j] = strconv.Itoa(count)
				} else {
					result[i][j] = "*"
				}
			}
		}
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				fmt.Print(result[i][j])
			}
			fmt.Println()
		}

	}

}
