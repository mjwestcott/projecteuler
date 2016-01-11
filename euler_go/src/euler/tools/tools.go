// tools.go
//
// This package implements several functions useful for solving project euler
// problems.

package tools

import (
	"math"
	"sort"
)

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

//-----------------------------------------------------------------------------
// Sorting

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int           { return len(s) }

// SortedString returns a new sorted string
func SortedString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

//-----------------------------------------------------------------------------
// Reversing

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
