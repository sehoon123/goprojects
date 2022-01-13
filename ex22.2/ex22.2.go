package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	e := q.v.Front()
	q.v.Remove(e)
	return e.Value
}

func NewQueue() *Queue {
	return &Queue{v: list.New()}
}

func main() {

	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
