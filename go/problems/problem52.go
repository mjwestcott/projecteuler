// problem52.go
//
// https://projecteuler.net/problem=52
//
// It can be seen that the number, 125874, and its double, 251748, contain exactly
// the same digits, but in a different order.
//
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
// contain the same digits.

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
)

// multiples returns [2x 3x 4x 5x 6x]
func multiples(x int) []int {
	var ms []int
	for i := 2; i < 7; i++ {
		ms = append(ms, x*i)
	}
	return ms
}

// sameDigits checks whether two ints contain the same digits.
func sameDigits(x, y int) bool {
	a := tools.SortedString(strconv.Itoa(x))
	b := tools.SortedString(strconv.Itoa(y))
	return a == b
}

// allSameDigits checks whether x has the same digits as all
// of 2x, 3x, 4x, 5x, and 6x.
func allSameDigits(x int) bool {
	for _, y := range multiples(x) {
		if sameDigits(x, y) == false {
			return false
		}
	}
	return true
}

// problem52 loops over ints from 1 to inf, returning the smallest x such that
// 2x, 3x, 4x, 5x, and 6x, contain the same digits.
func problem52() int {
	n := 1
	for {
		if allSameDigits(n) {
			return n
		}
		n++
	}
}

func main() {
	fmt.Println(problem52())
}
