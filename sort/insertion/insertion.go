package main

import "fmt"

// INSERTION SORT
func main() {
	array := []int{7, 5, 9, 0, 3, 1, 6, 2, 4, 8}

	fmt.Println(array)
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	fmt.Println(array)
}
