package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc        = bufio.NewScanner(os.Stdin)
	wr        = bufio.NewWriter(os.Stdout)
	parent    []int
	rank      []int
	score     []int
	tempScore []int
	tempCount []int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m, k := nextInt(), nextInt(), nextInt()
	parent = make([]int, n+1)
	rank = make([]int, n+1)
	score = make([]int, n+1)
	tempCount = make([]int, n+1)
	tempScore = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		parent[i] = i
		rank[i] = i
	}

	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		score[a] = b
		tempScore[a] = b
		tempCount[a] = 1
	}

	for i := 0; i < k; i++ {
		a, b := nextInt(), nextInt()
		union(a, b)
	}

	fmt.Fprintln(wr, parent)
	fmt.Fprintln(wr, tempScore)
	fmt.Fprintln(wr, tempCount)
	count := 0
	sum := 0.0
	for i := 1; i < n+1; i++ {
		if score[i] != 0 {
			sum += float64(score[i])
			count++
		} else if tempScore[parent[i]] != 0 {
			sum += math.Floor(float64(tempScore[parent[i]]) / float64(tempCount[parent[i]]))
			count++
		}
	}
	fmt.Fprintf(wr, "%.2f\n", sum/float64(count))
	//fmt.Fprintln(wr, sum/float64(count))
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	sx, sy := score[x], score[y]
	x = find(x)
	y = find(y)

	if x == y {
		return
	}
	if rank[x] < rank[y] {
		parent[y] = x
		if sy != 0 {
			tempScore[x] += sy
			tempCount[x]++
		}
	} else {
		parent[x] = y
		if sx != 0 {
			tempScore[y] += sx
			tempCount[y]++
		}
	}

	if rank[x] == rank[y] {
		rank[x]++
	}
}
