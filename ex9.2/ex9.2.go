package main

import "fmt"

func main() {
	var price int
	var member int
	var rich bool

	price = 35000
	rich = true
	member = 4

	if price > 50000 {
		if rich == true {
			fmt.Println("신발끈 묶기")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price < 50000 && price > 30000 {
		if member > 3 {
			fmt.Println("나눠내자")
		} else {
			fmt.Println("신발끈 묶기")
		}
	} else if price < 30000 {
		fmt.Println("내가 낸다")
	}

}
