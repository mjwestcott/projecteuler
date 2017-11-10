// problem58.go
//
// https://projecteuler.net/problem=58
//
// Starting with 1 and spiralling anticlockwise in the following way, a square
// spiral with side length 7 is formed.
//
//     37 36 35 34 33 32 31
//     38 17 16 15 14 13 30
//     39 18  5  4  3 12 29
//     40 19  6  1  2 11 28
//     41 20  7  8  9 10 27
//     42 21 22 23 24 25 26
//     43 44 45 46 47 48 49
//
// It is interesting to note that the odd squares lie along the bottom right
// diagonal, but what is more interesting is that 8 out of the 13 numbers lying
// along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.
//
// If one complete new layer is wrapped around the spiral above, a square
// spiral with side length 9 will be formed. If this process is continued, what
// is the side length of the square spiral for which the ratio of primes along
// both diagonals first falls below 10%?

package main

import (
	"euler/tools"
	"fmt"
	"math"
)

// Given the bottom right corner number n, return the length of the square spiral.
// Note: n will always be a square number 9, 25, 49, etc.
func squareLength(n int) int {
	return int(math.Sqrt(float64(n)))
}

// Given the bottom right corner number, return the four corner numbers.
// e.g. 49 -> [49 43 37 31]
func corners(n int) []int {
	x := squareLength(n) - 1
	return []int{n, n - x, n - (2 * x), n - (3 * x)}
}

// Return a new slice holding only the elements of xs that are prime.
func filterPrime(xs []int) []int {
	var ys []int
	for _, x := range xs {
		if tools.IsPrime(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Find the smallest square length for which primes / total < 0.1
func problem58() int {
	length := 7
	primes := 8
	total := 13
	for float64(primes)/float64(total) > 0.1 {
		length += 2
		primes += len(filterPrime(corners(length * length)))
		total += 4
	}
	return length
}

func main() {
	fmt.Println(problem58())
}
