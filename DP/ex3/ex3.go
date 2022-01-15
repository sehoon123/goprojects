package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	d := [1001]int{}
	d[1] = 1
	d[2] = 3

	for i := 3; i <= n; i++ {
		d[i] = (d[i-1] + d[i-2] + d[i-2]) % 796796
	}

	fmt.Println(d[n])

}
