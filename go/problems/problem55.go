// problem55.go
//
// https://projecteuler.net/problem=55
//
// If we take 47, reverse and add, 47 + 74 = 121, which is palindromic. A
// number that never forms a palindrome through the reverse and add process is
// called a Lychrel number. How many Lychrel numbers are there below
// ten-thousand? (Only consider fifty iterations)

package main

import (
	"euler/tools"
	"fmt"
	"math/big"
)

// isPalindrome checks whether the given *big.Int is a palindrome.
func isPalindrome(n *big.Int) bool {
	a := n.String()
	b := tools.ReversedString(a)
	return a == b
}

// reversed returns a new *big.Int with reversed digits.
func reversed(n *big.Int) *big.Int {
	a := n.String()
	b := tools.ReversedString(a)
	c, _ := new(big.Int).SetString(b, 10)
	return c
}

// isLychrel checks whether the given int is a Lychrel number.
func isLychrel(n int) bool {
	const limit = 50
	z := big.NewInt(int64(n))
	for i := 0; i < limit; i++ {
		z.Add(z, reversed(z)) // Set z to z + reversed(z)
		if isPalindrome(z) {
			return false
		}
	}
	return true
}

func problem55() int {
	count := 0 // How many numbers satisfy the isLychrel predicate?
	const limit = 10000
	for i := 0; i < limit; i++ {
		if isLychrel(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(problem55())
}
