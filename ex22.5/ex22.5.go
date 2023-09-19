package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"iPhone", 50000}
	m[46] = Product{"MacBook", 100000}
	m[78] = Product{"iPad", 30000}
	m[345] = Product{"MacBook Pro", 150000}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
