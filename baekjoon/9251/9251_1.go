package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var line1, line2 string
	fmt.Fscanln(r, &line1)
	fmt.Fscanln(r, &line2)

	arr := make([][]int, 26)
	for i, v := range line1 {
		arr[v-'A'] = append(arr[v-'A'], i)
	}
	fmt.Fprintln(w, arr)

	arr2 := make([]int, 0, len(line2))

	for _, v := range line2 {
		if arr[v-'A'] == nil {
			continue
		}
		if len(arr[v-'A']) != 1 {
			arr2 = append(arr2, arr[v-'A'][0])
			arr[v-'A'] = arr[v-'A'][1:]
		} else {
			arr2 = append(arr2, arr[v-'A'][0])
		}
	}
	fmt.Fprintln(w, arr2)

	d := make([]int, len(arr2))
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < i; j++ {
			if arr2[i] > arr2[j] && d[i] < d[j] {
				d[i] = d[j]
			}
		}
		d[i] += 1
	}

	sort.Ints(d)
	if len(d) == 0 {
		fmt.Fprintln(w, 0)
	} else {
		fmt.Fprintln(w, d[len(d)-1])
	}

}
