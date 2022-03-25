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

const (
	INF = int(1e9)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type edge struct {
	to, cost int
}

type PriorityQueue []edge

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(edge))
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
	v, e, p := nextInt(), nextInt(), nextInt()
	graph := make([][]edge, v+1)
	for i := 0; i < e; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		graph[a] = append(graph[a], edge{b, c})
		graph[b] = append(graph[b], edge{a, c})
	}

	//fmt.Fprintln(wr, graph)

	pq := PriorityQueue{}
	heap.Push(&pq, edge{1, 0})
	value := make([]int, v+1)
	for i := range value {
		value[i] = INF
	}
	value[1] = 0

	for len(pq) > 0 {
		cur := heap.Pop(&pq).(edge)

		if value[cur.to] < cur.cost {
			continue
		}

		for _, next := range graph[cur.to] {
			if value[next.to] >= value[cur.to]+next.cost {
				value[next.to] = value[cur.to] + next.cost
				heap.Push(&pq, edge{next.to, value[next.to]})
			}
		}
	}
	heap.Push(&pq, edge{p, 0})
	temp := make([]int, v+1)
	for i := range temp {
		temp[i] = INF
	}
	temp[p] = 0
	for len(pq) > 0 {
		cur := heap.Pop(&pq).(edge)

		if temp[cur.to] < cur.cost {
			continue
		}

		for _, next := range graph[cur.to] {
			if temp[next.to] >= temp[cur.to]+next.cost {
				temp[next.to] = temp[cur.to] + next.cost
				heap.Push(&pq, edge{next.to, temp[next.to]})
			}
		}
	}

	if value[v] == value[p]+temp[v] {
		fmt.Fprintln(wr, "SAVE HIM")
	} else {
		fmt.Fprintln(wr, "GOOD BYE")
	}

}
