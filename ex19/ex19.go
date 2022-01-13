package main

import(
	"fmt"
)

func addNum(slice []int) {
	slice = append(slice, 4)
}

func main() {
	slice := []int{1,2,3}
	addNum(slice)

	fmt.Println(slice)
}
