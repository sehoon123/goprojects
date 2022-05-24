package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	for {
		max := 0
		n, k := nextInt(), nextInt()
		temp := 1
		for i := 1; temp <= k || i <= n; i++ {
			temp *= i
			if temp%k == 0 || k%temp == 0 {
				max = temp
			}
		}
		fmt.Println(max)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//func lcm(a, b int) int {
//	return a * (b / gcd(a, b))
//}
