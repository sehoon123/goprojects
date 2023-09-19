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

	n := nextInt()
	arr := make([]int, n)
	visited1 := make([]bool, n)
	visited2 := make([]bool, n)
	visited3 := make([]bool, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	result1, result2, result3 := 0, 0, 0
	now1, now2, now3 := 0, 1, 2
	for {
		if !visited3[now3] {
			visited3[now3] = true
			now3 += arr[now3]
			result3++
		} else {
			break
		}
	}
	for {
		if !visited2[now2] {
			visited2[now2] = true
			now2 += arr[now2]
			result2++
		} else {
			break
		}
	}
	for {
		if !visited1[now1] {
			visited1[now1] = true
			now1 += arr[now1]
			result1++
		} else {
			break
		}
	}
	answer := max(result1, result2, result3)
	fmt.Fprintln(wr, answer+1)
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
