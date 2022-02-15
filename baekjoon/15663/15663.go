package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	wr     = bufio.NewWriter(os.Stdout)
	N, M   int
	nums   []int
	picked []int
	used   []bool
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	N, M = scanInt(), scanInt()

	nums = make([]int, N)
	for ni := 0; ni < N; ni++ {
		nums[ni] = scanInt()
	}

	sort.Ints(nums)
	picked = make([]int, M)
	used = make([]bool, N)

	pick(0)

	wr.Flush()
}

func scanInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func pick(idx int) {
	if idx == M {
		printAns()
		return
	}

	before := -1
	for i := 0; i < N; i++ {
		if nums[i] == before || used[i] {
			continue
		}

		used[i] = true
		picked[idx] = nums[i]
		before = nums[i]
		pick(idx + 1)
		used[i] = false
	}
}

func printAns() {
	for _, v := range picked {
		wr.WriteString(strconv.Itoa(v))
		wr.WriteByte(' ')
	}
	wr.WriteByte('\n')
}
