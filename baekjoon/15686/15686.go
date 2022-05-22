package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc                    = bufio.NewScanner(os.Stdin)
	wr                    = bufio.NewWriter(os.Stdout)
	graph, house, chicken [][]int
	visited               []bool
	arr                   []int
	distance              []int
	ans                   int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()

	graph = make([][]int, n)
	chicken = make([][]int, 0+m-m)
	house = make([][]int, 0)
	//distance = make([]int, n)

	// making graph
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = nextInt()
			// memorize chicken and house
			if graph[i][j] == 2 {
				chicken = append(chicken, []int{i, j})
			} else if graph[i][j] == 1 {
				house = append(house, []int{i, j})
			}
		}
	}

	visited = make([]bool, len(chicken)+1)
	arr = make([]int, 0)
	ans = int(1e9)
	dfs(len(chicken), m)

	fmt.Fprintln(wr, ans)

}

func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}

func dfs(n, m int) {
	if len(arr) == m {
		dummy := make([]int, 0)
		for i := 0; i < len(house); i++ {
			temp := int(1e9)
			for _, v := range arr {
				dist := int(math.Abs(float64(house[i][0]-chicken[v-1][0]))) + int(math.Abs(float64(house[i][1]-chicken[v-1][1])))
				if temp > dist {
					temp = dist
				}
			}
			dummy = append(dummy, temp)
		}
		if ans > sum(dummy) {
			ans = sum(dummy)
		}
		return
	}

	for i := 1; i < n+1; i++ {
		if !visited[i] {
			visited[i] = true
			arr = append(arr, i)
			dfs(n, m)
			arr = arr[:len(arr)-1]
			for j := i + 1; j < n+1; j++ {
				visited[j] = false
			}
		}
	}
}
