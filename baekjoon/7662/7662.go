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

func nextChar() string {
	sc.Scan()
	return sc.Text()
}

type Item struct {
	value  int
	maxIdx int
	minIdx int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].maxIdx = i
	h[j].maxIdx = j
}

func (h *MaxHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.maxIdx = n
	*h = append(*h, item)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.maxIdx = -1
	*h = old[0 : n-1]
	return item
}

type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].minIdx = i
	h[j].minIdx = j
}

func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.minIdx = n
	*h = append(*h, item)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.minIdx = -1
	*h = old[0 : n-1]
	return item
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	T := nextInt()
	for c := 0; c < T; c++ {
		k := nextInt()
		max := make(MaxHeap, 0, k)
		min := make(MinHeap, 0, k)
		for i := 0; i < k; i++ {
			command, num := nextChar(), nextInt()
			if command == "D" {
				if len(max) == 0 {
					continue
				}
				if num == 1 {
					item := heap.Pop(&max).(*Item)
					heap.Remove(&min, item.minIdx)
				} else if num == -1 {
					item := heap.Pop(&min).(*Item)
					heap.Remove(&max, item.maxIdx)
				}
			} else if command == "I" {
				item := &Item{value: num}
				heap.Push(&max, item)
				heap.Push(&min, item)
			}
		}
		if len(max) == 0 {
			fmt.Fprintln(wr, "EMPTY")
		} else {
			fmt.Fprintln(wr, max[0].value, min[0].value)
		}
	}
}
