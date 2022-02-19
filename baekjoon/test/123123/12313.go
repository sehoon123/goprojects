package main

import "fmt"

func main() {
	graph := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 0, 1, 1, 0, 1},
		{0, 0, 0, 0, 1, 0},
	}
	temp := make([][]int, len(graph))
	for i := 0; i < len(graph); i++ {
		temp[i] = make([]int, len(graph[i]))
		copy(temp[i], graph[i])
	}
	fmt.Println(graph)
	temp[0][0] = 1
	fmt.Println(&graph)
	fmt.Println(&temp)
}
