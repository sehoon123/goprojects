package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)

	for {
		var n int
		fmt.Scanln(&n)
		if n > number {
			fmt.Println("Too high")
		} else if n < number {
			fmt.Println("Too low")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}
