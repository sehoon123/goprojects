package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}

	fmt.Println("Sum of", a, "to", b, "is", sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 100000000)
	}
	wg.Wait()
}
