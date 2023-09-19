package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 1; i < n+1; i++ {
		arr[i-1] = i
	}

	for len(arr) > 1 {
		arr = arr[1:]
		arr = rotate(&arr)
	}
	fmt.Println(arr[0])

}

func rotate(arr *[]int) []int {
	return append((*arr)[1:len(*arr)], (*arr)[:1]...)
}
