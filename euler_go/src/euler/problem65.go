// problem65.go
//
// https://projecteuler.net/problem=65
//
// The infinite continued fraction can be written, √2 = [1;(2)], (2) indicates
// that 2 repeats ad infinitum. In a similar way, √23 = [4;(1,3,1,8)].
//
// It turns out that the sequence of partial values of continued fractions for
// square roots provide the best rational approximations. Let us consider the
// convergents for √2. The sequence of the first ten convergents for √2 are:
//
//     1, 3/2, 7/5, 17/12, 41/29, 99/70, 239/169, 577/408, 1393/985, 3363/2378, ...
//
// What is most surprising is that the important mathematical constant,
//
//     e = [2; 1,2,1, 1,4,1, 1,6,1 , ... , 1,2k,1, ...].
//
// The first ten terms in the sequence of convergents for e are: 2, 3, 8/3,
// 11/4, 19/7, 87/32, 106/39, 193/71, 1264/465, 1457/536, ...
//
// The sum of digits in the numerator of the 10th convergent is 1+4+5+7=17.
// Find the sum of digits in the numerator of the 100th convergent of the
// continued fraction for e.

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Return n values from the sequence 1,2,1, 1,4,1, 1,6,1 1,8,1, ...
// (The pattern in the continued fractional representation of e.)
func partialValues(n int) []int {
	var result []int
	x := 2
	for i := 0; i <= n/3; i++ {
		result = append(result, 1, x, 1)
		x += 2
	}
	return result[:n]
}

// Return the nth convergent of the continued fraction for e.
func e(n int) *big.Rat {
	two := big.NewRat(2, 1)
	if n == 1 {
		return two
	}
	// Collect the first n-1 partial values of e.
	values := partialValues(n - 1)
	// Construct the continued fraction, where 'tail' is the recursive component.
	one := big.NewRat(1, 1)
	return new(big.Rat).Add(two, new(big.Rat).Quo(one, tail(values)))
}

// Recursively return the tail end of the continued fractional representation of e.
func tail(values []int) *big.Rat {
	next, values := values[0], values[1:]
	x := big.NewRat(int64(next), 1)
	if len(values) == 0 {
		return x
	}
	one := big.NewRat(1, 1)
	return new(big.Rat).Add(x, new(big.Rat).Quo(one, tail(values)))
}

func problem65() int {
	digits := e(100).Num().String()
	sum := 0
	for _, d := range digits {
		x, _ := strconv.Atoi(string(d))
		sum += x
	}
	return sum
}

func main() {
	ans := problem65()
	fmt.Println(ans)
}
