package main

import "fmt"

//quick_sort
func quick_sort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quick_sort(arr, left, pivot-1)
		quick_sort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	array := []int{5, 7, 9, 0, 3, 1, 6, 2, 4, 8}
	quick_sort(array, 0, len(array)-1)
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
}
