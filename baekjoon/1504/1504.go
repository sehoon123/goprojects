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

type Node struct {
	next, distance int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N, E := nextInt(), nextInt()

	graph := make([][]Node, N+1)

	for i := 0; i < E; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		graph[a] = append(graph[a], Node{b, c})
	}

	v1, v2 := nextInt(), nextInt()

	result := min(dijkstra(1, v1, graph)+dijkstra(v1, v2, graph)+dijkstra(v2, N, graph), dijkstra(1, v2, graph)+dijkstra(v2, v1, graph)+dijkstra(v1, N, graph))
	fmt.Fprintln(wr, result)

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dijkstra(start, end int, graph [][]Node) int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = 1e9
	}
	dist[start] = 1
	pq := PriorityQueue{}
	heap.Push(&pq, &Node{start, 1})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if dist[node.next] > node.distance {
			continue
		}
		for _, n := range graph[node.next] {
			if n.distance+node.distance < dist[n.next] {
				dist[n.next] = n.distance + node.distance
				heap.Push(&pq, &Node{n.next, dist[n.next]})
			}
		}
	}

	return dist[end] - 1
}
