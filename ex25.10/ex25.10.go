package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	ch <- 3
	var a int
	var p *int
	p = &a

	*p = <-ch
	fmt.Println(a)

	fmt.Println(a)
}
