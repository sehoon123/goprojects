package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("one")
		break
	case 2 :
		fmt.Println("two")
		break
	case 3:
		fmt.Println("three")
		fallthrough
	case 4 :
		fmt.Println("four")
	default:
		fmt.Println("default")
	}
}
