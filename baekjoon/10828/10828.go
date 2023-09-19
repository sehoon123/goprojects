package main

<<<<<<< HEAD
import "fmt"

func main() {
	fmt.Println("vim-go")
=======
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n := nextInt()
	stack := make([]int, 0, 10000)

	for i := 0; i < n; i++ {
		command := nextWord()
		switch command {
		case "push":
			x := nextInt()
			stack = append(stack, x)
		case "pop":
			if len(stack) == 0 {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case "size":
			fmt.Fprintln(wr, len(stack))
		case "empty":
			if len(stack) == 0 {
				fmt.Fprintln(wr, 1)
			} else {
				fmt.Fprintln(wr, 0)
			}
		case "top":
			if len(stack) == 0 {
				fmt.Fprintln(wr, -1)
			} else {
				fmt.Fprintln(wr, stack[len(stack)-1])
			}
		}

	}
>>>>>>> 4736327f175ef60ddbe484ff18c9db872c26a72a
}
