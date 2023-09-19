package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}, {5, 6}, {6, 4}}
	visited := make([]bool, 7)
	dfs(graph, 1, &visited)

}

//dfs use Stack
func dfs(graph [][]int, v int, visited *[]bool) {
	(*visited)[v] = true
	fmt.Printf("%d ", v)
	for _, u := range graph[v] {
		if !(*visited)[u] {
			dfs(graph, u, visited)
		}
	}
}
