// problem77.go
//
// It is possible to write ten as the sum of primes in exactly five different
// ways:
//
//     7 + 3
//     5 + 5
//     5 + 3 + 2
//     3 + 3 + 2 + 2
//     2 + 2 + 2 + 2 + 2
//
// What is the first value which can be written as the sum of primes in over
// five thousand different ways?

package main

import (
	"euler/tools"
	"fmt"
)

func numPartitions(n int, primes []int) int {
	// Using a slightly different algorithm than problem 76.
	// This one is adapted from SICP: https://mitpress.mit.edu/sicp/full-text/book/book-Z-H-11.html
	// See the section entitled 'Example: Counting change'. Their logic is
	// more intuitive than that which I presented in the previous problem.
	switch {
	case n < 0:
		return 0
	case n == 0:
		return 1
	case len(primes) == 0:
		return 0
	default:
		return numPartitions(n, primes[1:]) + numPartitions(n-primes[0], primes)
	}
}

// Return a key sutiable for use in a map.
func k(n int, primes []int) string {
	return fmt.Sprintf("%d %v", n, primes)
}

var cache = make(map[string]int)

// A memoized version of numPartitions.
func memoNumPartitions(n int, primes []int) int {
	key := k(n, primes)
	if _, ok := cache[key]; !ok {
		cache[key] = numPartitions(n, primes)
	}
	return cache[key]

}

func problem77() int {
	var primes []int
	for i := 2; i < 100; i++ {
		if tools.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	// What is the first value which can be written as the sum of primes in
	// over five thousand different ways?
	for i := 2; ; i++ {
		if x := memoNumPartitions(i, primes); x > 5000 {
			return i
		}
	}
}

func main() {
	ans := problem77()
	fmt.Println(ans)
}
