// problem70.go
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used
// to determine the number of positive numbers less than or equal to n which
// are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all
// less than nine and relatively prime to nine, φ(9)=6. The number 1 is
// considered to be relatively prime to every positive number, so φ(1)=1.
//
// Interestingly, φ(87109)=79180, and it can be seen that 87109 is a
// permutation of 79180.
//
// Find the value of n, 1 < n < 10**7, for which φ(n) is a permutation of n and
// the ratio n/φ(n) produces a minimum.

package main

import (
	"euler/tools"
	"fmt"
	"math"
	"strconv"
)

func isPerm(x, y int) bool {
	a := strconv.Itoa(x)
	b := strconv.Itoa(y)
	return tools.SortedString(a) == tools.SortedString(b)
}

// The search space is too large for brute-force. So, note that we are seeking
// roughly the inverse of the previous problem -- to minimize n/phi(n).
// Therefore, we want to maximize phi(n), which is acheived for numbers with
// the fewest and largest unique prime factors. But the number cannot simply be
// prime because in that case phi(n) == n-1 which is not a permutation of n.
// Therefore, the best candidates should have two unique prime factors.
func problem70() int {
	// Since we are seeking large values for both prime factors, we can
	// search among numbers close to the value of sqrt(1e7) ~ 3162
	var primes []int
	for i := 2000; i < 4000; i++ {
		if tools.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	// Keep track of the minimal n/phi(n) ratio found
	// and the value of n that produced it.
	min := math.Inf(1)
	res := 0

	const limit = 10000000
	for _, x := range primes {
		for _, y := range primes {
			if n := x * y; n < limit && x != y {
				if pn := tools.Phi(n); isPerm(n, pn) {
					if rat := float64(n) / float64(pn); rat < min {
						min = rat
						res = n
					}
				}
			}
		}
	}
	return res
}

func main() {
	ans := problem70()
	fmt.Println(ans)
}
