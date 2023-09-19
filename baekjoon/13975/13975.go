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

	T := nextInt()
	for i := 0; i < T; i++ {
		n := nextInt()
		arr := make(PriorityQueue, 0, n)
		for j := 0; j < n; j++ {
			heap.Push(&arr, nextInt())
		}

		ans := 0
		for len(arr) > 1 {
			a, b := heap.Pop(&arr), heap.Pop(&arr)
			ans += a.(int) + b.(int)
			heap.Push(&arr, a.(int)+b.(int))
		}
		fmt.Fprintln(wr, ans)
	}

}

type PriorityQueue []int

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
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
