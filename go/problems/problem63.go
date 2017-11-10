// problem63.go
//
// https://projecteuler.net/problem=63
//
// The 5-digit number, 16807=7**5, is also a fifth power. Similarly, the
// 9-digit number, 134217728=8**9, is a ninth power. How many n-digit positive
// integers exist which are also an nth power?

package main

import (
	"euler/tools"
	"fmt"
	"math/big"
)

func pow(x, y int) *big.Int {
	a := new(big.Int)
	b, c := big.NewInt(int64(x)), big.NewInt(int64(y))
	return a.Exp(b, c, nil)
}

// Return the list of powers to which one can raise n such that the result of
// exponentiation is an integer with number of digits == power
// findPowers(6) -> [1 2 3 4]
func findPowers(x int) []int {
	var ys []int
	for y := 1; tools.NumDigits(pow(x, y)) == y; y++ {
		ys = append(ys, y)
	}
	return ys
}

// Find the sum of the lengths of all the results returned by findPowers(i).
// When findPowers(i) returns a nil slice, it indicates no more results.
func problem63() int {
	sum := 0
	for i := 1; findPowers(i) != nil; i++ {
		sum += len(findPowers(i))
	}
	return sum
}

func main() {
	fmt.Println(problem63())
}
