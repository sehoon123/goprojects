package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	cargo := make([]int, n)
	for i, v := range cargo {
		fmt.Scan(&v)
		cargo[i] = v
	}
	fmt.Println(cargo)

	d[0] = cargo[0]
	d[1] = int(math.Max(float64(cargo[0]), float64(cargo[1])))

	for i := 2; i < n; i++ {
		d[i] = int(math.Max(float64(d[i-1]), float64(d[i-2]+cargo[i])))
	}

	fmt.Println(d[n-1])

}
