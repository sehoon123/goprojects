package main

import "fmt"

type House struct {
	Address string
	Size int
	Price float64
	Category string
}

func main() {
	var house House
	house.Address = "123 Main St"
	house.Size = 1000
	house.Price = 1000000
	house.Category = "Condo"

	fmt.Printf("%v\n", house)
	fmt.Println(house)
}
