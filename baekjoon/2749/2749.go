package main

import (
	"fmt"
)

var (
	fib map[int][]int
	MOD = 1000000
)

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	if n <= 2 {
		fmt.Println(1)
		return
	}

	fib = make(map[int][]int)
	fib[1] = []int{1, 1, 1, 0}

	ans := fibonacci(n)
	fmt.Println(ans[1])
}

func fibonacci(x int) []int {
	_, ok := fib[x]
	if ok {
		return fib[x]
	}

	a := x / 2
	b := x - a

	_, ok = fib[a]
	if !ok {
		fib[a] = fibonacci(a)
	}

	_, ok = fib[b]
	if !ok {
		fib[b] = fibonacci(b)
	}

	fib[x] = mulMatrix(fib[a], fib[b])

	return fib[x]
}

func mulMatrix(a, b []int) []int {
	res := make([]int, 4)
	res[0] = (a[0]*b[0] + a[1]*b[2]) % MOD
	res[1] = (a[0]*b[1] + a[1]*b[3]) % MOD
	res[2] = (a[2]*b[0] + a[3]*b[2]) % MOD
	res[3] = (a[2]*b[1] + a[3]*b[3]) % MOD
	return res
}
