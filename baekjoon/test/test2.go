package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	var wg sync.WaitGroup
	wg.Add(len(slice))

	fmt.Println("running")

	for i := 0; i < len(slice); i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(slice[i])
		}(i)
	}

	wg.Wait()

}
