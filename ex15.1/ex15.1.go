package main

import "fmt"

func main() {
	poet1 := "죽는 날 까지 하늘을 우러러\n 한 점 부끄럼이 없기를"
	poet2 := `죽는 날 까지 하늘을 우러러
	훈점 부끄럼이 없기를,
	잎새에 이는 바람\n에도
	나는 괴로워했다.`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
