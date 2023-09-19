package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	schedules := make([][]int, n)
	for i := 0; i < n; i++ {
		schedules[i] = make([]int, 2)
		fmt.Fscanf(r, "%d %d\n", &schedules[i][0], &schedules[i][1])
	}

	sort.Slice(schedules, func(i, j int) bool {
		if schedules[i][1] == schedules[j][1] {
			return schedules[i][0] < schedules[j][0]
		}
		return schedules[i][1] < schedules[j][1]
	})

	result := make([][]int, 0, n)
	result = append(result, schedules[0])

	for i := 1; i < n; i++ {
		if schedules[i][0] >= result[len(result)-1][1] {
			result = append(result, schedules[i])
		}
	}
	fmt.Fprintln(w, len(result))
}
