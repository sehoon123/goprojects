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

type queue struct {
	arr []int
}

func (q *queue) push(x int) {
	q.arr = append(q.arr, x)
}

func (q *queue) pop() int {
	if len(q.arr) == 0 {
		return -1
	}
	x := q.arr[0]
	q.arr = q.arr[1:]
	return x
}

func (q *queue) size() int {
	return len(q.arr)
}

func (q *queue) empty() bool {
	return len(q.arr) == 0
}

func (q *queue) front() int {
	if q.empty() {
		return -1
	}
	return q.arr[0]
}

func (q *queue) back() int {
	if q.empty() {
		return -1
	}
	return q.arr[len(q.arr)-1]
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	q := queue{}
	n := nextInt()
	for i := 0; i < n; i++ {
		command := nextWord()
		switch command {
		case "push":
			q.push(nextInt())
		case "pop":
			wr.WriteString(strconv.Itoa(q.pop()) + "\n")
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
