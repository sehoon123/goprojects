package main

import "fmt"

func main() {
	const pi1 float64 = 3.14159265358979323846
	var pi2 = 3.14159265358979323846

	//pi1 = 3
	pi2 = 4

	fmt.Printf("pi1 = %f\n", pi1)
	fmt.Printf("pi2 = %f\n", pi2)
}
