package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 1024*1024)
	sc.Scan()
	str := sc.Text()
	fmt.Println(len(str))

}
