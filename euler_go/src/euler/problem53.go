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
	x := factorial(n)
	y := factorial(r)
	z := factorial(n - r)
	denom := new(big.Int).Mul(y, z)
	return new(big.Int).Div(x, denom)
}

func problem53() int {
	// Initialize target as one-million.
	var target big.Int
	target.Exp(big.NewInt(10), big.NewInt(6), nil)

	sum := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			nCr := numCombinations(n, r)
			if nCr.Cmp(&target) == 1 { // nCr > target
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
