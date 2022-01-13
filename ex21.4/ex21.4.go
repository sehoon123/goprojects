package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func getOperator(op string) func(int, int) int {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "*":
		return func (a, b int) int {
			return a * b
		}
	default:
		panic("unsupported operator")
	}
}

func main() {
	var Operator = getOperator("+")
	k := Operator(1, 2)
	fmt.Println(k)
}