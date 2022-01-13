package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice1 = append(slice1[:idx], append([]int{100}, slice1[idx:]...)...)
	//for i := idx+1; i < len(slice1); i++ {
	//	slice1[i-1] = slice1[i]
	//}

	//slice1 = slice1[:len(slice1)-1]

	fmt.Println(slice1, cap(slice1))
	sort.Ints(slice1)
	fmt.Println(slice1, cap(slice1))

}
