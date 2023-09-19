package main

import "fmt"

func main() {
	//nums := [...]int{10,20,30,40,50}
	//
	//nums[2] = 300

	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(nums[i])
	//}

	var t = [5]float64{24.0, 25.9, 27.8, 26.9, 28.9}

	for i, v := range t {
		fmt.Println(i, v)
	}
}
