package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	n     int
	graph [][]int
)

func nextInt() []int {
	sc.Scan()
	temp := make([]int, n)
	for i, c := range sc.Bytes() {
		temp[i] = int(c - '0')
	}
	return temp
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	fmt.Scan(&n)
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = nextInt()
	}
	f(0, 0, n, n)
}

func f(sx, sy, ex, ey int) {
	flag := true
	for i := sx; i < ex; i++ {
		for j := sy; j < ey; j++ {
			if graph[i][j] != graph[sx][sy] {
				flag = false
				fmt.Fprint(wr, "(")
				f(sx, sy, sx+(ex-sx)/2, sy+(ey-sy)/2)
				f(sx, sy+(ey-sy)/2, sx+(ey-sy)/2, ey)
				f(sx+(ey-sy)/2, sy, ex, sy+(ey-sy)/2)
				f(sx+(ey-sy)/2, sy+(ey-sy)/2, ex, ey)
				fmt.Fprint(wr, ")")
				return
			}
		}
	}
	if flag {
		if graph[sx][sy] == 0 {
			fmt.Fprint(wr, "0")
		} else {
			fmt.Fprint(wr, "1")
		}
	}

}
