package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	wr   = bufio.NewWriter(os.Stdout)
	arr  []int
	temp []int
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
	temp = make([]int, 0, n)
	arr = make([]int, 0, n)
	cnt := 0

	for i := 0; i < n; i++ {
		a := nextInt()
		if a == 1 {
			b := nextInt()
			idx := sort.Search(len(arr), func(i int) bool { return arr[i] >= b })
			arr = append(arr, 0)
			copy(arr[idx+1:], arr[idx:])
			arr[idx] = b
			cnt++
		} else if a == 2 {
			b, c := nextInt(), nextInt()
			result := sort.Search(len(arr), func(i int) bool { return arr[i] > b })
			if result < c {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, arr[result-c])
			}
		} else if a == 3 {
			b, c := nextInt(), nextInt()
			result := sort.Search(len(arr), func(i int) bool { return arr[i] >= b })
			if result+(c-1) > cnt-1 {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, arr[result+(c-1)])
			}
		}

	}

}
