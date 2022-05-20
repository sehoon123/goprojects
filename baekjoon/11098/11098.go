package main

import (
	"bufio"
	"os"
	"sort"
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

func nextWord() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	T := nextInt()

	for i := 0; i < T; i++ {
		N := nextInt()
		list := make(PlayerList, N)
		for j := 0; j < N; j++ {
			list[j].price = nextInt()
			list[j].name = nextWord()
		}

		sort.Sort(list)
		wr.WriteString(list[0].name + "\n")
	}
}

type Player struct {
	price int
	name  string
}

type PlayerList []Player

func (p PlayerList) Len() int {
	return len(p)
}

func (p PlayerList) Less(i, j int) bool {
	return p[i].price > p[j].price
}

func (p PlayerList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
