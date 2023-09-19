package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type message struct {
	city, cost int
}

type messages []*message

func (m messages) Len() int { return len(m) }
func (m messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m messages) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}
func (m *messages) Push(x interface{}) {
	*m = append(*m, x.(*message))
}
func (m *messages) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, C int
	fmt.Fscan(r, &N, &M, &C)
	graph := make([][]message, N+1)

	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		graph[a] = append(graph[a], message{b, c})
	}

	value := make([]int, N+1)
	pq := messages{}
	heap.Push(&pq, &message{C, 0})
	value[C] = 0

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*message)

		if value[now.city] < now.cost {
			continue
		}

		for _, next := range graph[now.city] {
			if value[next.city] > next.cost || value[next.city] == 0 {
				value[next.city] = next.cost
				heap.Push(&pq, &message{next.city, next.cost})
			}
		}
	}

	fmt.Fprintln(w, value)
}
