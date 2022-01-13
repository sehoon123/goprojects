package main

import "fmt"

func main() {
	a :=  [5]int{1, 2, 3, 4, 5}
	b := [5]int{100, 200, 300, 400, 500}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
