package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc       = bufio.NewScanner(os.Stdin)
	wr       = bufio.NewWriter(os.Stdout)
	arr      [][]int16
	inDegree []int16
	n, m     int16
)

func nextInt() int16 {
	sc.Scan()
	r := 0
	for _, c := range sc.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return int16(r)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m = nextInt(), nextInt()

	inDegree = make([]int16, n+1)
	arr = make([][]int16, n+1)

	var prev int16
	for i := 0; i < int(m); i++ {
		tmp := nextInt()
		for j := 0; j < int(tmp); j++ {
			now := nextInt()
			if j != 0 {
				arr[prev] = append(arr[prev], now)
				inDegree[now]++
			}
			prev = now
		}
	}

	topologySort()
}

func topologySort() {
	result := make([]int16, 0, n)
	queue := make([]int16, 0, n)

	for i := 1; i <= int(n); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, int16(i))
		}
	}

	for i := 1; i <= int(n); i++ {
		if len(queue) == 0 {
			fmt.Fprintln(wr, 0)
			return
		}
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)
		for j := 0; j < len(arr[v]); j++ {
			u := arr[v][j]
			inDegree[u]--
			if inDegree[u] == 0 {
				queue = append(queue, u)
			}
		}

	}

	for i := 1; i <= int(n); i++ {
		fmt.Fprintf(wr, "%d\n", result[i-1])
	}

}
