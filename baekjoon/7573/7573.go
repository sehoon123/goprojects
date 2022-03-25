package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	n, l, m := nextInt(), nextInt(), nextInt()
	fish := make([][2]int, n-n+m)
	for i := 0; i < m; i++ {
		fish[i][0], fish[i][1] = nextInt()-1, nextInt()-1
	}

	max := 0
	for i := range fish {
		for j := 1; j < l/2; j++ {
			nx := j
			ny := l/2 - j
			for k := fish[i][0] - nx; k <= fish[i][0]; k++ {
				for w := fish[i][1] - ny; w <= fish[i][1]; w++ {
					count := 0
					for v := range fish {
						if k <= fish[v][0] && fish[v][0] <= k+nx && w <= fish[v][1] && fish[v][1] <= w+ny {
							count++
						}
					}
					if count > max {
						max = count
					}
				}
			}
		}
	}
	wr.WriteString(strconv.Itoa(max) + "\n")
}
