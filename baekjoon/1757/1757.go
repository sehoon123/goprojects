package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	n, m  int
	arr   []int
	cache [][][]int
)

//FAST IO
func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m = nextInt(), nextInt()
	arr = make([]int, n)
	cache = make([][][]int, n+1)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
		cache[i+1] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			cache[i+1][j] = make([]int, 2)
		}
	}

	fmt.Fprintln(wr, run(1, 0, 0))
}

//status 0 : continue, 1 : break
func run(idx, gauge, status int) int {
	if idx == n+1 {
		if gauge == 0 {
			return 0
		}
		//마지막에 게이지가 0이 아니면 최대가 될수 없게 안전장치
		return -1000000000
	}
	// 캐시에 있으면 바로 리턴
	if cache[idx][gauge][status] != 0 {
		return cache[idx][gauge][status]
	}
	//휴식중이면 0될때까지 계속 휴식
	if status == 1 {
		// 게이지가 0됐을때
		if gauge == 0 {
			cache[idx][gauge][status] = max(run(idx+1, gauge+1, 0)+arr[idx-1], run(idx+1, 0, 1))
			// 휴식중인데 게이지가 0이 아닐때
		} else {
			cache[idx][gauge][status] = run(idx+1, gauge-1, 1)
		}
		//달리는 중일때
	} else {
		// 게이지가 다 찼으면 휴식
		if gauge+1 > m {
			cache[idx][gauge][status] = run(idx+1, gauge-1, 1)
			// 게이지가 0일때 index error 방지용
		} else if gauge-1 < 0 {
			cache[idx][gauge][status] = max(run(idx+1, gauge+1, 0)+arr[idx-1], run(idx+1, 0, 1))
		} else {
			cache[idx][gauge][status] = max(run(idx+1, gauge+1, 0)+arr[idx-1], run(idx+1, gauge-1, 1))
		}
	}
	return cache[idx][gauge][status]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
