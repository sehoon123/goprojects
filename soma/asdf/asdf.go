package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var line string
	fmt.Scan(&line)
	char := []rune(line)
	result := make([]rune, 0, 100000)
	for _, v := range char {
		//fmt.Println(v - 'a')
		for k := 0; k < 26; k++ {
			fmt.Println((a*k + b) % 26)
			fmt.Println(v - 'a')
			if rune((a*k+b)%26) == v-'a' {
				fmt.Println(k)
				//fmt.Println(((int(v-'a') + 26*k) - b) / a)
				result = append(result, rune(k))
				break
			}
		}
	}

	for _, v := range result {
		fmt.Print(string(v + 'a'))
	}

}
