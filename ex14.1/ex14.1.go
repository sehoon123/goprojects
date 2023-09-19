package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var p1 = &a
	var p2 = &a
	var p3 = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p1 == p3: %v\n", p1 == p3)

}
