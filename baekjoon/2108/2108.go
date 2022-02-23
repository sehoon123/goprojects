package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	wr  = bufio.NewWriter(os.Stdout)
	n   int
	arr []int
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n = nextInt()
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}
	fmt.Fprintln(wr, mean(arr))
	fmt.Fprintln(wr, median(arr))
	fmt.Fprintln(wr, Mode(arr))
	fmt.Fprintln(wr, scope(arr))
}

func mean(arr []int) int {
	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	return int(math.Round(sum / float64(len(arr))))
}

func median(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)/2]
}

func Mode(arr []int) int {
	cntArr := make(map[int]int, len(arr))
	for _, v := range arr {
		cntArr[v]++
	}
	max := 0
	for _, v := range cntArr {
		if v > max {
			max = v
		}
	}
	temp := make([]int, 0, len(cntArr))

	for i, v := range cntArr {
		if v == max {
			temp = append(temp, i)
		}
	}
	sort.Ints(temp)
	if len(temp) == 1 {
		return temp[0]
	} else {
		return temp[1]
	}
}

func scope(arr []int) int {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}
