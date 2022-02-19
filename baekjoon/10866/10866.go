package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type deque struct {
	arr []int
}

func (q *deque) push_front(x int) {
	q.arr = append(q.arr, 0)
	copy(q.arr[1:], q.arr)
	q.arr[0] = x
}

func (q *deque) push_back(x int) {
	q.arr = append(q.arr, x)
}

func (q *deque) pop_front() int {
	if len(q.arr) == 0 {
		return -1
	}
	x := q.arr[0]
	q.arr = q.arr[1:]
	return x
}
func (q *deque) pop_back() int {
	if len(q.arr) == 0 {
		return -1
	}
	x := q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	return x
}

func (q *deque) size() int {
	return len(q.arr)
}

func (q *deque) empty() bool {
	return len(q.arr) == 0
}

func (q *deque) front() int {
	if q.empty() {
		return -1
	}
	return q.arr[0]
}

func (q *deque) back() int {
	if q.empty() {
		return -1
	}
	return q.arr[len(q.arr)-1]
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	q := deque{}
	n := nextInt()
	for i := 0; i < n; i++ {
		command := nextWord()
		switch command {
		case "push_front":
			q.push_front(nextInt())
		case "push_back":
			q.push_back(nextInt())
		case "pop_front":
			wr.WriteString(strconv.Itoa(q.pop_front()) + "\n")
		case "pop_back":
			wr.WriteString(strconv.Itoa(q.pop_back()) + "\n")
		case "size":
			wr.WriteString(strconv.Itoa(q.size()) + "\n")
		case "empty":
			if q.empty() {
				wr.WriteString("1\n")
			} else {
				wr.WriteString("0\n")
			}
		case "front":
			wr.WriteString(strconv.Itoa(q.front()) + "\n")
		case "back":
			wr.WriteString(strconv.Itoa(q.back()) + "\n")
		}
	}
}
