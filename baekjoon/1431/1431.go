package main

import (
	"bufio"
	"os"
	"regexp"
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

type serial struct {
	num string
}

type serial_arr []serial

func (s serial_arr) Len() int {
	return len(s)
}

func (s serial_arr) Less(i, j int) bool {
	if len(s[i].num) < len(s[j].num) {
		return true
	} else if len(s[i].num) == len(s[j].num) {
		re := regexp.MustCompile("[0-9]+")
		matched1 := re.FindAllString(s[i].num, -1)
		matched2 := re.FindAllString(s[j].num, -1)
		result1 := 0
		result2 := 0
		for _, v := range matched1 {
			for k, _ := range v {
				result1 += int(v[k] - '0')
			}
		}
		for _, v := range matched2 {
			for k, _ := range v {
				result2 += int(v[k] - '0')
			}
		}
		//fmt.Println("result1:", result1, "result2:", result2)
		if result1 < result2 {
			return true
		} else if result1 == result2 {
			return s[i].num < s[j].num
		} else {
			return false
		}
	}
	return false
}

func (s serial_arr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	arr := make(serial_arr, n)
	for i := 0; i < n; i++ {
		arr[i].num = nextWord()
	}
	sort.Sort(arr)

	for _, v := range arr {
		wr.WriteString(v.num + "\n")
	}
}
