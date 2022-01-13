package main

import (
	"errors"
	"fmt"
)

//func division(x, y int) (int, int) {
//	if y == 0 {
//		panic("y cannot be zero")
//	}
//	return x / y, x % y
//}

//error는 계속 실행함
func division(x, y int) error {
	if y == 0 {
		return errors.New("y cannot be zero")
	}
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	return nil
}

func main() {
	division(10, 0)
	division(10, 2)
}
