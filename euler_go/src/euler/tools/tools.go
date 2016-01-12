// tools.go
//
// This package implements several functions useful for solving project euler
// problems.

package tools

import (
	"math"
	"sort"
)

// Min returns the minimum of two int values.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

//-----------------------------------------------------------------------------
// Sorting

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }

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

// ReversedString returns a new reversed string
func ReversedString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
