package main

import "fmt"

func main() {
	d := [3]int{1, 2, 3}
	f := d[:]
	fmt.Println(d)
	fmt.Println(f)
	f[2] = 10
	fmt.Println(d)
	fmt.Println(f)
}
