package main

import (
	"fmt"
	"time"
)

func main() {
	var t1 = time.Timer{}
	var t2 = time.NewTimer(time.Second)
	fmt.Println(t1)
	fmt.Println(t2)
}
