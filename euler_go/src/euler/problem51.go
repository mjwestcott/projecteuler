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
	"math"
)

// toDigits converts an int to a []int representing its digits.
// e.g. 1234 -> [1 2 3 4]
func toDigits(n int) []int {
	var ds []int
	for n != 0 {
		ds = append(ds, n%10)
		n = int(n / 10)
	}
	tools.ReverseInts(ds)
	return ds
}

// toNum converts a []int representing digits to an int
// e.g. [1 2 3 4] -> 1234
func toNum(digits []int) int {
	var n int
	for i, x := range tools.ReversedInts(digits) {
		n += x * int(math.Pow(10, float64(i)))
	}
	return n
}

// Our strategy is as follows. Since we are seeking an eight prime family, it
// must be the case that the pattern of digits which are replaced contains
// either 0, 1, or 2 in the smallest family member. Therefore, we can search
// through primes and replace digits in patterns specified by the locations 0,
// 1, and 2. If the family of numbers that results contains eight primes, we
// have found the solution.

// In the example given, 56003 is the smallest member of an eight prime family.
// We would find the pattern of 0s at indices (2, 3) to produce the
// corresponding family from 56**3.

// findIndices returns three slices, where each contains the indices in the
// given number of the digits 0, 1, and 2 respectively.
// e.g. 18209912 -> [[3], [0 6], [2 7]]
// e.g. 56003 -> [[2 3], [], []]
func findIndices(n int) [][]int {
	var indices [][]int
	for _, target := range []int{0, 1, 2} {
		var found []int
		for i, x := range toDigits(n) {
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
	var fam []int
	digits := toDigits(n)
	for i := 0; i < 10; i++ {
		for _, idx := range indices {
			digits[idx] = i
		}
		// return sentinel value (-1) in case of leading zero
		if digits[0] == 0 {
			fam = append(fam, -1)
		} else {
			fam = append(fam, toNum(digits))
		}
	}
	return fam
}

// isSmallestMember checks whether the given number satisfies
// the problem description.
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

func problem51() int {
	for p := range tools.GetPrimesFrom(56995) {
		if isSmallestMember(p) {
			return p
		}
	}
	return -1
}

func main() {
	ans := problem51()
	fmt.Println(ans)
}
