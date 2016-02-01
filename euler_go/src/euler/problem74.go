// problem74.go
//
// The number 145 is well known for the property that the sum of the factorial
// of its digits is equal to 145:
//
//     1! + 4! + 5! = 1 + 24 + 120 = 145
//
// Perhaps less well known is 169, in that it produces the longest chain of
// numbers that link back to 169; it turns out that there are only three such
// loops that exist:
//
//     169 → 363601 → 1454 → 169
//     871 → 45361 → 871
//     872 → 45362 → 872
//
// It is not difficult to prove that EVERY starting number will eventually get
// stuck in a loop. For example,
//
//     69 → 363600 → 1454 → 169 → 363601 (→ 1454)
//     78 → 45360 → 871 → 45361 (→ 871)
//     540 → 145 (→ 145)
//
// Starting with 69 produces a chain of five non-repeating terms, but the
// longest non-repeating chain with a starting number below one million is
// sixty terms. How many chains, with a starting number below one million,
// contain exactly sixty non-repeating terms?

package main

import (
	"euler/tools"
	"fmt"
)

func factorial(n int) int {
	// We're only dealing with factorial(x) where x < 10
	// in this problem, so returning an int is fine.
	if n == 0 {
		return 1
	}
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
}

// To memoize the sumFactorialDigits function.
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if _, ok := cache[n]; !ok {
			val := f(n)
			cache[n] = val
		}
		return cache[n]
	}
}

// Return the sum of the factorial of every digit in n.
var sumFactorialDigits = memoize(func(n int) int {
	digits := tools.IntToDigits(n)
	sum := 0
	for _, d := range digits {
		sum += factorial(d)
	}
	return sum
})

func contains(xs []int, y int) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func problem74() int {
	// Known chain loop lengths given in problem description. We will use
	// this map to cache all further results as we calculate them.
	knownLoops := map[int]int{145: 1, 169: 3, 1454: 3, 871: 2, 872: 2, 69: 5, 78: 4, 540: 2}
	for i := 1; i <= 1000000; i++ {
		chain := []int{i}
		next := sumFactorialDigits(i)
		for {
			if contains(chain, next) {
				// We have found a new loop, add to the cache.
				knownLoops[i] = len(chain)
				break
			}
			if val, ok := knownLoops[next]; ok {
				// We have found a known loop, add its length to the current chain.
				knownLoops[i] = len(chain) + val
				break
			}
			// We haven't found a loop, continue to investigate the chain.
			chain = append(chain, next)
			next = sumFactorialDigits(next)
		}
	}
	count := 0
	for _, x := range knownLoops {
		if x == 60 {
			count++
		}
	}
	return count
}

func main() {
	ans := problem74()
	fmt.Println(ans)
}
