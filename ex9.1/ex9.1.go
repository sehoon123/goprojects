package main

import "fmt"

var cnt	int

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}
func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 increased")
	}

}
