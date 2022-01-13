package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}, {5, 6}, {6, 4}}
	visited := make([]bool, 7)
	bfs(graph, 1, &visited)

}

func bfs(graph [][]int, start int, visited *[]bool) {
	(*visited)[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)
		for _, neighbor := range graph[node] {
			if !(*visited)[neighbor] {
				(*visited)[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

}
