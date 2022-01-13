package main

import "fmt"

func main() {
	t := [...]int{10, 20, 30}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}
