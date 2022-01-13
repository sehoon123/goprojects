package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter text: ")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Error: ", err)

			//키보드 버퍼 지우기
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("%d\n", number)
		if number % 2 == 0 {
			break
		}
	}

	fmt.Println("Done")
}
