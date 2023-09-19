package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a, b, A, B := nextInt(), nextInt(), nextInt(), nextInt()
	if A < a {
		a, b, A, B = A, B, a, b
	}

	c, d, C, D := nextInt(), nextInt(), nextInt(), nextInt()
	if C < c {
		c, d, C, D = C, D, c, d
	}

	if a == A {
		if B < b {
			b, B = B, b
		}
		if c <= a && a <= C {
			if (b <= d && d <= B) || (b <= D && D <= B) {
				fmt.Fprintln(wr, 1)
				return
			} else {
				fmt.Fprintln(wr, 0)
				return
			}
		} else {
			fmt.Fprintln(wr, 0)
			return
		}
	}

	if c == C {
		if D < d {
			d, D = D, d
		}
		if a <= c && c <= A {
			if (b <= d && d <= B) || (b <= D && D <= B) {
				fmt.Fprintln(wr, 1)
				return
			} else {
				fmt.Fprintln(wr, 0)
				return
			}
		} else {
			fmt.Fprintln(wr, 0)
			return
		}
	}

	m1 := float64(B-b) / float64(A-a)
	m2 := float64(D-d) / float64(C-c)

	if m1 == m2 {
		if B < b {
			b, B = B, b
		}
		if float64(b)-m1*float64(a) == float64(d)-m2*float64(c) {
			if (b <= d && d <= B) || (b <= D && D <= B) || (d <= b && b <= D) || (d <= B && B <= D) {
				fmt.Fprintln(wr, 1)
				return
			} else {
				fmt.Fprintln(wr, 0)
				return
			}
		} else {
			fmt.Fprintln(wr, 0)
			return
		}
	}

	res := (m1*float64(a) - m2*float64(c) + float64(d) - float64(b)) / (m1 - m2)
	res = math.Floor(res*10000000000) / 10000000000

	if (float64(a) <= res && res <= float64(A)) && (float64(c) <= res && res <= float64(C)) {
		fmt.Fprintln(wr, 1)
	} else {
		fmt.Fprintln(wr, 0)
	}
}
