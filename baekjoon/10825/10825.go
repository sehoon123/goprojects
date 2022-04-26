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
	students := make(studentSlice, n)
	for i := 0; i < n; i++ {
		student := Student{
			name:    nextWord(),
			korean:  nextInt(),
			english: nextInt(),
			math:    nextInt(),
		}
		students = append(students, student)
	}

	sort.Sort(students)
	for _, student := range students {
		wr.WriteString(student.name + "\n")
	}

}

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

type studentSlice []Student

func (s studentSlice) Len() int {
	return len(s)
}
func (s studentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s studentSlice) Less(i, j int) bool {
	if s[i].korean > s[j].korean {
		return true
	} else if s[i].korean == s[j].korean {
		if s[i].english < s[j].english {
			return true
		} else if s[i].english == s[j].english {
			if s[i].math > s[j].math {
				return true
			} else if s[i].math == s[j].math {
				return s[i].name < s[j].name
			}
		}
	}
	return false
}
