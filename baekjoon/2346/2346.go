package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		arr[i][0] = i + 1
		arr[i][1] = nextInt()
	}
	kkk := arr
	for i := 0; i < n; i++ {
		idx, num := kkk[0][0], kkk[0][1]
		kkk = remove(kkk, 0)
		fmt.Fprintf(wr, "%d ", idx)
		if len(kkk) == 0 {
			break
		}

		if num > 0 {
			kkk = rotate(kkk, num-1)
		} else {
			kkk = rotate(kkk, len(kkk)+num%len(kkk))
		}
	}

}
func rotate(arr [][]int, k int) [][]int {
	k = k % len(arr)
	arr = append(arr[k:], arr[:k]...)
	return arr
}

func remove(arr [][]int, k int) [][]int {
	return append(arr[:k], arr[k+1:]...)
}
