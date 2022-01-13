package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func getOperator(op string) func (int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return multiply
	} else {
		return nil
	}
}

func main() {
	var operator func (int, int) int
	operator = getOperator("+")

	var result = operator(2, 3)
	fmt.Println(result)
}