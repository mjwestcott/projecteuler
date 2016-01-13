// problem51.go
//
// By replacing the 1st digit of the 2-digit number *3, it turns out that six of
// the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.
//
// By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit
// number is the first example having seven primes among the ten generated
// numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and
// 56993. Consequently 56003, being the first member of this family, is the
// smallest prime with this property.
//
// Find the smallest prime which, by replacing part of the number (not necessarily
// adjacent digits) with the same digit, is part of an eight prime value family.

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
)

// Our strategy is as follows. Since we are seeking an eight prime family, it
// must be the case that the pattern of digits which are replaced contains
// either 0, 1, or 2 in the smallest family member. Therefore, we can search
// through primes and replace digits in patterns specified by the locations 0,
// 1, and 2. If the family of numbers that results contains eight primes, we
// have found the solution.
//
// In the example given, 56003 is the smallest member of an eight prime family.
// We would find the pattern of 0s at indices (2, 3) to produce the
// corresponding family from 56**3.

// findIndices returns three slices, where each contains the indices in the
// given number of the digits 0, 1, and 2 respectively.
// e.g. 18209912 -> [[3], [0 6], [2 7]]
// e.g. 56003 -> [[2 3], [], []]
func findIndices(n int) [][]int {
	var indices [][]int
	for _, target := range "012" {
		var found []int
		for i, x := range strconv.Itoa(n) {
			if x == target {
				found = append(found, i)
			}
		}
		indices = append(indices, found)
	}
	return indices
}

// family returns the family of numbers resulting from replacing
// digits at the specific indices with the digits 0 to 9.
// e.g. 56003, [2 3] -> [56003, 56113, 56223, 56333, 56443, ...]
func family(n int, indices []int) []int {
	var ms []int
	template := strconv.Itoa(n)
	for _, x := range "0123456789" {
		// Build new family members by replacing characters
		// and then converting to an int.
		member := []rune(template)
		for _, i := range indices {
			member[i] = x
		}
		// Return sentinel value (-1) in case of leading zero.
		if member[0] == '0' {
			ms = append(ms, -1)
		} else {
			m, _ := strconv.Atoi(string(member))
			ms = append(ms, m)
		}
	}
	return ms
}

// isSmallestMember checks whether the given number satisfies the problem
// description: is it the smallest member of an 8-prime family?
func isSmallestMember(n int) bool {
	for _, indices := range findIndices(n) {
		var primes []int
		for _, x := range family(n, indices) {
			if tools.IsPrime(x) {
				primes = append(primes, x)
			}
		}
		if len(primes) == 8 {
			return true
		}
	}
	return false
}

// problem51 loops over ints from 56995 to inf, returning the smallest member
// of an 8-prime family.
func problem51() int {
	n := 56995 // given in the description as a member of a 7-prime family
	for {
		if tools.IsPrime(n) && isSmallestMember(n) {
			return n
		}
		n++
	}
}

func main() {
	ans := problem51()
	fmt.Println(ans)
}
