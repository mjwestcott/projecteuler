// problem53.go
//
// https://projecteuler.net/problem=53
//
// How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are
// greater than one-million?

package main

import (
	"euler/tools"
	"fmt"
	"math/big"
)

func numCombinations(n, r int) *big.Int {
	x := tools.Factorial(n)
	y := tools.Factorial(r)
	z := tools.Factorial(n - r)
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
	fmt.Println(problem53())
}
