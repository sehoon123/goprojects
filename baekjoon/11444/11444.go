package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	wr  = bufio.NewWriter(os.Stdout)
	mat map[int][]int
)

const (
	MOD = 1000000007
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()

	if n <= 2 {
		if n == 0 {
			fmt.Fprintln(wr, 0)
			return
		}
		fmt.Fprintln(wr, 1)
		return
	}

	mat = make(map[int][]int, 61)
	mat[1] = []int{1, 1, 1, 0}

	ans := fibo(n)
	fmt.Fprintln(wr, ans[1])
}

func fibo(x int) []int {
	if mat[x] != nil {
		return mat[x]
	}

	a := x / 2
	b := x - a

	if mat[a] == nil {
		mat[a] = fibo(a)
	}

	if mat[b] == nil {
		mat[b] = fibo(b)
	}

	mat[x] = matMul(mat[a], mat[b])

	return mat[x]
}

func matMul(a, b []int) []int {
	result := make([]int, 4)
	result[0] = (a[0]*b[0] + a[1]*b[2]) % MOD
	result[1] = (a[0]*b[1] + a[1]*b[3]) % MOD
	result[2] = (a[2]*b[0] + a[3]*b[2]) % MOD
	result[3] = (a[2]*b[1] + a[3]*b[3]) % MOD
	return result
}
