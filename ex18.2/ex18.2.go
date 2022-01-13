package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		println("empty")
	}

	slice = append(slice,1)
	fmt.Println(slice)
}