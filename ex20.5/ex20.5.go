package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int :
		fmt.Printf("%d is an int\n", t)
	case float64:
		fmt.Printf("%f is a float64\n", t)
	case string:
		fmt.Printf("%s is a string\n", t)
	default:
		fmt.Printf("%v is a unknown type\n", t)
	}

}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(10.5)
	PrintVal("hello")
	PrintVal(Student{10})
}
