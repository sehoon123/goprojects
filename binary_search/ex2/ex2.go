package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	array := []int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		array = append(array, a)
	}
	sort.Ints(array)

	result := binary_search(array, array[0], array[n-1], m)
	fmt.Println(result)
}

func binary_search(array []int, start, end, target int) int {
	mid := (start + end) / 2
	sum := 0

	for _, v := range array {
		if v > mid {
			sum += v - mid
		}
	}

	if sum == target {
		return mid
	} else if sum > target {
		return binary_search(array, mid+1, end, target)
	} else {
		return binary_search(array, start, mid-1, target)
	}

}
