package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := make([]int, 0, 1000001)
	i := 0
	for len(arr) < 1000001 {
		check := make([]int, 100)
		i++
		temp := strconv.Itoa(i)
		flag := true
		for _, v := range temp {
			if check[int(v)] == 1 {
				flag = false
				break
			}
			check[int(v)] += 1
		}
		if flag {
			arr = append(arr, i)
		}
	}

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		fmt.Println(arr[n-1])
	}
}
