package main

import (
	"fmt"
	"math"
)

func main() {
	n := 27

	d := [30001]int{}
	//find minimum

	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + 1
		if i%5 == 0 {
			d[i] = int(math.Min(float64(d[i]), float64(d[i/5]+1)))
		}
		if i%3 == 0 {
			d[i] = int(math.Min(float64(d[i]), float64(d[i/3]+1)))
		}
		if i%2 == 0 {
			d[i] = int(math.Min(float64(d[i]), float64(d[i/2]+1)))
		}
	}

	fmt.Println(d[n])
}
