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

type edge struct {
	to, cost int
}

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, v := range sc.Bytes() {
		if v == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(v - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	V, E, K := nextInt(), nextInt(), nextInt()

	graph := make([][]edge, 20001)
	for i := 0; i < E; i++ {
		u, v, w := nextInt(), nextInt(), nextInt()
		graph[u] = append(graph[u], edge{v, w})
	}
	pq := PriorityQueue{}
	value := make([]int, 20001)
	heap.Push(&pq, &edge{K, 0})
	value[K] = 1

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*edge)
		if value[now.to] < now.cost {
			continue
		}

		for _, next := range graph[now.to] {
			if value[next.to] == 0 || value[next.to] > value[now.to]+next.to {
				value[next.to] = value[now.to] + next.cost
				heap.Push(&pq, &edge{next.to, value[next.to]})
			}
		}
	}

	for i := 1; i <= V; i++ {
		if value[i] == 0 {
			fmt.Fprintln(wr, "INF")
		} else {
			fmt.Fprintln(wr, value[i]-1)
		}
	}
}

type PriorityQueue []*edge

func (pq PriorityQueue) Len() int      { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*edge))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
