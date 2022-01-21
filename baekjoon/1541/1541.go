package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	sc.Split(bufio.ScanWords)
	sc.Scan()

	str := strings.Split(sc.Text(), "")
	str = append(str, "")

	var tmp string
	var flag int
	var sum int

	for i := 0; i < len(str); i++ {
		if str[i] == "-" || str[i] == "+" || i == len(str)-1 {

			if flag == 1 {
				aaa, _ := strconv.Atoi(tmp)
				sum -= aaa
			} else {
				aaa, _ := strconv.Atoi(tmp)
				sum += aaa
			}

			if str[i] == "-" {
				flag = 1
			}

			tmp = ""
			continue
		}

		tmp += str[i]
	}

	fmt.Fprintln(w, sum)

}
