package main

import (
	"bufio"
	"container/heap"
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
	n, m := nextInt(), nextInt()
	graph := make([][]edge, n+1)

	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		graph[a] = append(graph[a], edge{b, c})
		graph[b] = append(graph[b], edge{a, c})
	}

	pq := PriorityQueue{}
	fmt.Fprintln(wr, graph)
	value := make([]int, n+1)
	heap.Push(&pq, &edge{1, 0})
	value[1] = 0

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*edge)

		if value[now.to] < now.cost {
			continue
		}

		for _, next := range graph[now.to] {
			if value[next.to] == 0 || value[next.to] >= value[now.to]+next.cost {
				value[next.to] = value[now.to] + next.cost
				heap.Push(&pq, &edge{next.to, value[next.to]})
			}
		}
	}

	fmt.Fprintln(wr, value[n])

}

type edge struct {
	to, cost int
}

type PriorityQueue []*edge

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*edge))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[0 : n-1]
	x := old[n-1]
	return x
}
