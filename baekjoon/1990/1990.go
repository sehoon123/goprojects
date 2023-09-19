package main

import (
	"bufio"
	"fmt"
	"math"
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

	a, b := nextInt(), nextInt()
	//fmt.Fprintln(wr, []rune("123"))
	if b > 10000000 {
		b = 10000000
	}
	nums := make([]int, b+1)
	for i := 2; i < b+1; i++ {
		nums[i] = 1
	}

	for i := 2; i < int(math.Sqrt(float64(b+1)))+1; i++ {
		if nums[i] == 1 {
			j := 2
			for i*j < (b + 1) {
				nums[i*j] = 0
				j++
			}
		}
	}

	//fmt.Fprintln(wr, nums)
	for i := a; i <= b; i++ {
		if nums[i] == 1 {
			temp := strconv.Itoa(i)
			arr := []rune(temp)
			for {
				if len(arr) == 0 || len(arr) == 1 {
					fmt.Fprintln(wr, i)
					break
				}
				if arr[0] != arr[len(arr)-1] {
					break
				}

				arr = arr[1:]
				arr = arr[:len(arr)-1]
			}
		}
	}
	fmt.Fprintln(wr, -1)
}
