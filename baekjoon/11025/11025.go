package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	result := 0
	for i := 1; i <= n; i++ {
		result = (result + k) % i
	}
	fmt.Println(result + 1)

	//arr := make([]int, n)
	//for i := 0; i < n; i++ {
	//	arr[i] = i + 1
	//}
	//
	//for len(arr) > 1 {
	//	arr = rotate(arr, k-1)
	//	arr = arr[1:]
	//}
	//fmt.Println(arr[0])

}

//func rotate(arr []int, k int) []int {
//	n := len(arr)
//	k = k % n
//	arr = append(arr[k:], arr[0:k]...)
//	return arr
//}
