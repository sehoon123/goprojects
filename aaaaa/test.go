package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] > 5
	})
	fmt.Println(i)
}
