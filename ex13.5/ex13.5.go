package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}

func main() {
	var a = 8
	user := User{
		Age:   18,
		Score: 100.0,
	}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))
}
