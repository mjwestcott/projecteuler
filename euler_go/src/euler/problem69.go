// problem69.go
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used
// to determine the number of numbers less than n which are relatively prime to
// n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and
// relatively prime to nine, φ(9)=6.
//
// It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10. Find the value
// of n ≤ 1,000,000 for which n/φ(n) is a maximum.

package main

import (
	"euler/tools"
	"fmt"
)

// Note that the phi function multiplies n by (1 - (1/p)) for every p in
// its unique prime factors. Therefore, phi(n) will diminish as n has a greater
// number of small unique prime factors. Since we are seeking the largest value
// for n/phi(n), we want to minimize phi(n). We are therefore looking for the
// largest number <= 1e6 which is the product of the smallest unique prime
// factors, i.e successive prime numbers starting from 2.
func problem69() int {
	limit := 1000000
	candidate := 1 // The multiplicative identity.
	for i := 2; i*candidate <= limit; i++ {
		if tools.IsPrime(i) {
			candidate *= i
		}
	}
	return candidate
}

func main() {
	ans := problem69()
	fmt.Println(ans)
}
