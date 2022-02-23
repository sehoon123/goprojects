package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	wr   = bufio.NewWriter(os.Stdout)
	n, m int
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m = nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	target := make([]int, m)
	for i := 0; i < m; i++ {
		target[i] = nextInt()
	}

	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < len(a); j++ {
			if a[0] == target[i] {
				a = a[1:]
				break
			} else {
				if find(a, target[i]) < len(a)-find(a, target[i]) {
					for a[0] != target[i] {
						a = rotate_right(a)
						cnt++
					}
					a = a[1:]
					break
				} else {
					for a[0] != target[i] {
						a = rotate_left(a)
						cnt++
					}
					a = a[1:]
					break
				}
			}
		}
	}
	fmt.Fprintln(wr, cnt)

}

func rotate_right(arr []int) []int {
	return append(arr[1:], arr[:1]...)
}
func rotate_left(arr []int) []int {
	return append(arr[len(arr)-1:], arr[:len(arr)-1]...)
}

func find(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
