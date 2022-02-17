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
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Heap struct {
	arr []int
}

func (h Heap) Len() int {
	return len(h.arr)
}

func (h Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h Heap) Less(i, j int) bool {
	if abs(h.arr[i]) == abs(h.arr[j]) {
		return h.arr[i] < h.arr[j]
	}
	return abs(h.arr[i]) < abs(h.arr[j])
}

func (h *Heap) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *Heap) Pop() interface{} {
	n := h.Len()
	x := h.arr[n-1]
	h.arr = h.arr[:n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	h := &Heap{}
	heap.Init(h)
	n := nextInt()
	for i := 0; i < n; i++ {
		ele := nextInt()
		if ele == 0 {
			if h.Len() > 0 {
				wr.WriteString(strconv.Itoa(heap.Pop(h).(int)) + "\n")
			} else {
				wr.WriteString("0\n")
			}
		} else {
			heap.Push(h, ele)
		}
	}
}
