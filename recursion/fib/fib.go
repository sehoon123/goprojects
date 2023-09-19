package main

import "fmt"

var (
	cache []int
)

func main() {
	var n int
	fmt.Scan(&n)
	cache = make([]int, n+1)
	fmt.Println(fib(100))
}

func fib(n int) int {
	if cache[n] != 0 {
		return cache[n]
	}

	if n <= 1 {
		return 1
	}

	cache[n] = fib(n-1) + fib(n-2)
	return cache[n]
}
