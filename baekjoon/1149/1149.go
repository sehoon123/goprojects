package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	house := make([][3]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Fscan(r, &house[i][j])
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			temp := 0
			if j == 0 {
				temp = int(math.Min(float64(house[i-1][2]), float64(house[i-1][1])))
			} else if j == 1 {
				temp = int(math.Min(float64(house[i-1][0]), float64(house[i-1][2])))
			} else {
				temp = int(math.Min(float64(house[i-1][1]), float64(house[i-1][0])))
			}
			house[i][j] += temp
		}
	}

	sort.Ints(house[n-1][:])
	fmt.Println(house[n-1][0])

}
