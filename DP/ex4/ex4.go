package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	moneyList := make([]int, 0, n)
	d := [10001]int{}

	for i := 0; i < n; i++ {
		var money int
		fmt.Scanf("%d\n", &money)
		moneyList = append(moneyList, money)
	}

	for _, v := range moneyList {
		d[v]++
	}

	for i := 0; i < m+1; i++ {
		pivot := 10001
		for _, v := range moneyList {
			if i-v >= 0 && d[i-v] != 0 {
				pivot = int(math.Min(float64(pivot), float64(d[i-v]+d[v])))
				d[i] = pivot
			}
		}
	}

	fmt.Println(d[m])

}
