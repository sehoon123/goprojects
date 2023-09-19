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

	n := nextInt()
	list := make(PriorityQueue, 0)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		heap.Push(&list, []int{a, b})
	}

	time := make(PriorityQueue2, 0)

	for list.Len() > 0 {
		t := heap.Pop(&list).([]int)

		if len(time) == 0 {
			heap.Push(&time, t)
			continue
		}

		s := heap.Pop(&time).([]int)
		if t[0] >= s[1] {
			heap.Push(&time, t)
		} else {
			heap.Push(&time, s)
			heap.Push(&time, t)
		}
	}

	fmt.Fprintln(wr, len(time))

}

type PriorityQueue [][]int
type PriorityQueue2 [][]int

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}
func (pq *PriorityQueue2) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue2) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue2) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
