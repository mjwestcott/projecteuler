// tools.go
//
// This package implements several functions useful for solving project euler
// problems.

package tools

import "math"

// ReverseInts reverses a []int in place
func ReverseInts(xs []int) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

// ReversedInts returns a new reversed []int
func ReversedInts(xs []int) []int {
	res := make([]int, len(xs))
	copy(res, xs)
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// IsPrime checks whether n is prime
func IsPrime(n int) bool {
	if n < 3 {
		return n == 2
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GetPrimesFrom yields primes starting from the given value
func GetPrimesFrom(start int) <-chan int {
	ch := make(chan int)
	go func() {
		x := start
		for {
			if IsPrime(x) {
				ch <- x
			}
			x++
		}
	}()
	return ch
}
