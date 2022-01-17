package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	var a string
	fmt.Fscan(r, &a)

	sum := big.NewInt(0)
	aaaa := big.NewInt(1234567891)
	for i, v := range a {
		var s = big.NewInt(31)
		var t = big.NewInt(int64(i))
		alphabet := big.NewInt(int64(v - 'a' + 1))
		exp := s.Exp(s, t, nil)
		tmp := alphabet.Mul(alphabet, exp)
		tmp.Mod(tmp, aaaa)
		sum.Add(sum, tmp)
	}

	fmt.Fprintln(w, sum.Mod(sum, aaaa))
}
