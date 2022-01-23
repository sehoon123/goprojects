package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var k, n int
	fmt.Fscanf(r, "%d %d\n", &k, &n)

	list := make([]int, k)
	total := 0
	for i := range list {
		fmt.Fscan(r, &list[i])
		total += list[i]
	}
	alpha := total / n

	result := binary_search(list, alpha, n)
	fmt.Fprintln(w, result)

}

func binary_search(arr []int, alpha, target int) int {
	start, end := 1, alpha
	for start <= end {
		mid := (start + end) / 2
		beta := 0
		for _, v := range arr {
			beta += v / mid
		}
		if beta >= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}

//upperBound
func upperBound(arr []int, target int) int {
	start, end := 0, len(arr)
	for start < end {
		mid := (start + end) / 2
		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

//lowerBound
func lowerBound(arr []int, target int) int {
	start, end := 0, len(arr)
	for start < end {
		mid := (start + end) / 2
		if arr[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
