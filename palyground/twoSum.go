package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	//sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	//sc.Scan()
	//n, _ = strconv.Atoi(sc.Text())

	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		//sc.Scan()
		//a[i], _ = strconv.Atoi(sc.Text())
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &b[i])
		//sc.Scan()
		//a[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	var result = 0
	for i := 0; i < n; i++ {
		result += a[i] * b[i]
	}

	fmt.Println(result)
}
