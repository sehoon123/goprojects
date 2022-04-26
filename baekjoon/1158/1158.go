package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	fmt.Printf("%s", "<")
	for len(arr) > 1 {
		arr = rotate(arr, k-1)
		fmt.Printf("%d, ", arr[0])
		arr = arr[1:]
	}
	fmt.Printf("%d>\n", arr[0])
}

func rotate(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	arr = append(arr[k:], arr[0:k]...)
	return arr
}
