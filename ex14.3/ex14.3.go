package main

import "fmt"

type Data struct {
	value int
	data [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
	fmt.Printf("data[100] = %d\n", arg.data[100])
}

func main() {
	var data Data
	var p2 = new(Data)
	fmt.Println(p2)

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}