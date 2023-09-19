package main

import (
	"fmt"
	"math"
)

func fibonacci1(n float64) float64 {
	return ((math.Pow((1+math.Sqrt(5))/2, n)) - (math.Pow((1-math.Sqrt(5))/2, n))) / math.Sqrt(5)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci2(n-1) + fibonacci2(n-2)
}

func fibonacci3(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	one := 1
	two := 0
	rst := 0

	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func main() {
	fmt.Println(fibonacci3(30))
	fmt.Println(fibonacci1(30))
	fmt.Println(fibonacci2(30))
}
