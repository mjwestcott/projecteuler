// problem53.go
//
// How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are
// greater than one-million?

package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	x := new(big.Int)
	x.MulRange(1, int64(n))
	return x
}

func numCombinations(n, r int) *big.Int {
	a := factorial(n)
	b := factorial(r)
	c := factorial(n - r)
	denominator := b.Mul(b, c)
	result := a.Div(a, denominator)
	return result
}

func problem53() int {
	// Initialize limit as one-million.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(6), nil)

	sum := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			nCr := numCombinations(n, r)
			if nCr.Cmp(&limit) == 1 { // nCr > limit
				sum++
			}
		}
	}
	return sum
}

func main() {
	ans := problem53()
	fmt.Println(ans)
}
