package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	arr := make([][2]int, n+1)

	for i := 1; i < n+1; i++ {
		arr[i][0] = nextInt()
		arr[i][1] = nextInt()
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	left, right := arr[1][0], arr[1][1]
	ans := 0

	for i := 1; i < n+1; i++ {
		if right >= arr[i][1] {
			continue
		} else if arr[i][0] <= right && right < arr[i][1] {
			right = arr[i][1]
		} else {
			ans += right - left
			left, right = arr[i][0], arr[i][1]
		}
	}

	ans += right - left
	fmt.Fprintln(wr, ans)

}
