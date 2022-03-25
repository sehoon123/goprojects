package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc        = bufio.NewScanner(os.Stdin)
	wr        = bufio.NewWriter(os.Stdout)
	n         int
	graph     [][]int
	minus_one int
	zero      int
	one       int
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
	n = nextInt()
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = nextInt()
		}
	}

	f(0, 0, n, n)
	fmt.Fprintln(wr, minus_one, zero, one)

}

func f(sx, sy, ex, ey int) {
	flag := true
	for i := sx; i < ex; i++ {
		for j := sy; j < ey; j++ {
			if graph[i][j] != graph[sx][sy] {
				flag = false
				f(sx, sy, sx+(ex-sx)/3, sy+(ey-sy)/3)
				f(sx, sy+(ey-sy)/3, sx+(ex-sx)/3, sy+(ey-sy)/3*2)
				f(sx, sy+(ey-sy)/3*2, sx+(ex-sx)/3, ey)
				f(sx+(ex-sx)/3, sy, sx+(ex-sx)/3*2, sy+(ey-sy)/3)
				f(sx+(ex-sx)/3, sy+(ey-sy)/3, sx+(ex-sx)/3*2, sy+(ey-sy)/3*2)
				f(sx+(ex-sx)/3, sy+(ey-sy)/3*2, sx+(ex-sx)/3*2, ey)
				f(sx+(ex-sx)/3*2, sy, ex, sy+(ey-sy)/3)
				f(sx+(ex-sx)/3*2, sy+(ey-sy)/3, ex, sy+(ey-sy)/3*2)
				f(sx+(ex-sx)/3*2, sy+(ey-sy)/3*2, ex, ey)
				return
			}
		}
	}
	if flag {
		if graph[sx][sy] == 0 {
			zero++
		} else if graph[sx][sy] == 1 {
			one++
		} else {
			minus_one++
		}
	}
}
