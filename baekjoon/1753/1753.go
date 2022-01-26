package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

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
	V, E := nextInt(), nextInt()
	K := nextInt()

	graph := make([][]edge, V+1)
	for i := 0; i < E; i++ {
		u, v, w := nextInt(), nextInt(), nextInt()
		graph[u] = append(graph[u], edge{v, w})
	}

	pq := PriorityQueue{}
	value := make([]int, V+1)
	heap.Push(&pq, &edge{K, 0})
	value[K] = 1

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*edge)

		if value[now.to] < now.cost {
			continue
		}

		for _, next := range graph[now.to] {
			if value[next.to] == 0 || value[next.to] > value[now.to]+next.cost {
				value[next.to] = value[now.to] + next.cost
				heap.Push(&pq, &edge{next.to, value[next.to]})
			}
		}
	}

	for i := 1; i <= V; i++ {
		if value[i] == 0 {
			wr.WriteString("INF\n")
		} else {
			wr.WriteString(strconv.Itoa(value[i]-1) + "\n")
		}
	}

}

type edge struct {
	to, cost int
}

type PriorityQueue []*edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

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
