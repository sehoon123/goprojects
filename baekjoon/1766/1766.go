package main

import (
	"bufio"
	"container/heap"
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

	n, m := nextInt(), nextInt()
	graph := make([][]int, n+1)
	temp := make([]int, n+1)
	pq := make(PriorityQueue, 0, m)

	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		graph[a] = append(graph[a], b)
		temp[b]++
	}

	for i := 1; i <= n; i++ {
		if temp[i] == 0 {
			heap.Push(&pq, i)
		}
	}

	for len(pq) > 0 {
		u := heap.Pop(&pq).(int)
		fmt.Fprintf(wr, "%d ", u)
		for _, v := range graph[u] {
			temp[v]--
			if temp[v] == 0 {
				heap.Push(&pq, v)
			}
		}
	}

}

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
