package main

import "fmt"

func devide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a/ b, true
}

func main() {
	c, flag := devide(10, 2)
	fmt.Println(c, flag)
	d, flag := devide(9, 0)
	fmt.Println(d, flag)
}
