// problem72.go
//
// https://projecteuler.net/problem=72
//
// Consider the fraction, n/d, where n and d are positive integers. If n<d and
// HCF(n,d)=1, it is called a reduced proper fraction.
//
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order
// of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 21 elements in this set. How many elements
// would be contained in the set of reduced proper fractions for d ≤ 1,000,000?

package main

import "fmt"

// See https://en.wikipedia.org/wiki/Farey_sequence
// for the connection to Euler's phi function.
// In particular, note that |F n| = |F n-1| + phi(n).
func problem72() float64 {
	const limit = 1000000

	// By convention phi(0)=1 and phi(1)=1, but Project Euler
	// appears to disagree, so those values are set to zero.
	var phi [limit + 1]float64
	for i := 2; i <= limit; i++ {
		phi[i] = float64(i)
	}

	// Solved in the style of the sieve of Eratosthenes.
	marked := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		if !marked[i] {
			for j := i; j <= limit; j += i {
				phi[j] *= 1.0 - 1.0/float64(i)
				marked[j] = true
			}
		}
	}

	sum := 0.0
	for _, x := range phi {
		sum += x
	}
	return sum
}

func main() {
	ans := problem72()
	fmt.Println(ans)
}
