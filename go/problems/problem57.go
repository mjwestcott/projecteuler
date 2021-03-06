// problem57.go
//
// https://projecteuler.net/problem=57
//
// It is possible to show that the square root of two can be expressed as an
// infinite continued fraction.
//
//     √ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...
//
// By expanding this for the first four iterations, we get:
//
//     1 + 1/2 = 3/2 = 1.5
//     1 + 1/(2 + 1/2) = 7/5 = 1.4
//     1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
//     1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...
//
// The next three expansions are 99/70, 239/169, and 577/408, but the eighth
// expansion, 1393/985, is the first example where the number of digits in the
// numerator exceeds the number of digits in the denominator. In the first
// one-thousand expansions, how many fractions contain a numerator with more
// digits than denominator?

package main

import (
	"euler/tools"
	"fmt"
	"math/big"
)

func checkNumerator(x *big.Rat) bool {
	num, denom := x.Num(), x.Denom()
	return tools.NumDigits(num) > tools.NumDigits(denom)
}

// How many fractions contain a numerator with more digits than the denominator?
func problem57() int {
	sum := 0           // Number of fractions meeting the description.
	const limit = 1000 // Given in problem description.
	one := new(big.Rat).SetInt64(1)
	two := new(big.Rat).SetInt64(2)

	// result will be re-used each iteration to store the
	// current value of the fractional expansion.
	result := new(big.Rat)
	// tail will be re-used each iteration to store the
	// current value of the repeating component of the expansion.
	// That component is 2, (2 + 1/2), (2 + 1/(2 + 1/2)), ...
	tail := new(big.Rat).SetInt64(2)

	for i := 0; i < limit; i++ {
		temp := new(big.Rat)
		tail.Add(two, temp.Inv(tail))   // tail = (2 + 1/tail)
		result.Add(one, temp.Inv(tail)) // result = (1 + 1/tail)
		if checkNumerator(result) {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println(problem57())
}
