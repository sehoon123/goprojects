package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(gcd(n, m))
	fmt.Println(lcm(n, m))
}

func gcd(n, m int) int {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}

func lcm(n, m int) int {
	return n * m / gcd(n, m)
}
