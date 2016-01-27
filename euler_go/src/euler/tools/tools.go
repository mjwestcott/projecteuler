// tools.go
//
// This package implements several functions useful for solving project euler
// problems.

package tools

import (
	"math"
	"sort"
)

// Min returns the minimum of any number of int values.
func Min(xs ...int) int {
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// Max returns the maximum of any number of int values.
func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

// Sum returns the sum of any number of int values.
func Sum(xs ...int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
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

// Pow returns the integer power a**b.
func Pow(a, b int) int {
	// https://groups.google.com/d/msg/golang-nuts/PnLnr4bc9Wo/z9ZGv2DYxXoJ
	// Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Permutations returns all r-length permutations in lexicographic sort order.
// So, if the input iterable is sorted, the permutation tuples will be produced
// in sorted order.
func Permutations(xs []int, r int) [][]int {
	// Translated from the Python itertools module, see C source below for comments.
	// https://github.com/python/cpython/blob/master/Modules/itertoolsmodule.c#L3127
	pool := append([]int{}, xs...)
	n := len(pool)

	switch {
	case n == 0:
		panic("Permutations: passed zero-length slice")
	case r < 1 || r > n:
		panic("Permutations: passed bad r value")
	}

	indices := make([]int, n)
	cycles := make([]int, r)

	for i := 0; i < n; i++ {
		indices[i] = i
	}
	for i := 0; i < r; i++ {
		cycles[i] = n - i
	}

	var res [][]int
	res = append(res, pool[:r])

	for {
		var i int // Represents the leftmost element to change per iteration.
		for i = r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]
				perm := make([]int, r)
				for i, index := range indices[:r] {
					perm[i] = pool[index]
				}
				res = append(res, perm)
				break
			}
		}
		if i < 0 {
			return res
		}
	}
}

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
