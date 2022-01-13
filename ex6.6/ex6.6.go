package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}
func main() {
	var a = 0.1
	var b = 0.2
	var c = 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

}
