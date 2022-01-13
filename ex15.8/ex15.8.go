package main

import "fmt"

func main() {
	str := "Hello, 월드"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("TYPE: %T, STRING_VALUE: %c, VALUE: %d\n", arr[i],arr[i], arr[i])
	}
}
