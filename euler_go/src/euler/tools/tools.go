// tools.go
//
// This package implements several functions useful for solving project euler
// problems.

package tools

import (
	"math"
	"sort"
	"strconv"
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

// Unique returns the unique ints from a slice in order of first seen.
func Unique(xs []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			result = append(result, x)
		}
	}
	return result
}

// GCD returns the greatest common divisor of two ints.
func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// PrimeFactors returns all prime factors of n in ascending order.
func PrimeFactors(n int) []int {
	return pfactors(n, 2)
}

func pfactors(n, start int) []int {
	for i := start; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return append([]int{i}, pfactors(n/i, i)...)
		}
	}
	return []int{n}
}

// Phi is Euler's phi function (also known as Euler's totient function).
func Phi(n int) int {
	ps := Unique(PrimeFactors(n))
	x := float64(n)
	for _, p := range ps {
		x *= (1.0 - (1.0 / float64(p)))
	}
	return int(x)
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

// StringsToInts converts a slice of strings representing numbers
// to a slice of ints.
func StringsToInts(ss []string) []int {
	var xs []int
	for _, s := range ss {
		x, _ := strconv.Atoi(string(s))
		xs = append(xs, x)
	}
	return xs
}

// IntsToStrings converts a slice of ints to Itoa-converted strings.
func IntsToStrings(xs []int) []string {
	var ss []string
	for _, x := range xs {
		ss = append(ss, strconv.Itoa(x))
	}
	return ss
}

// IntToDigits converts an int to a slice of ints representing its digits.
func IntToDigits(n int) []int {
	var xs []int
	for n > 0 {
		xs = append(xs, n%10)
		n /= 10
	}
	ReverseInts(xs)
	return xs
}

// Stack is a simple generic stack implementation.
type Stack []interface{}

// Pop item off the stack. Panics if s is empty.
func (s *Stack) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	// Shrink the underlying array if the slice length <= 1/4 its capacity.
	if len(*s) <= cap(*s)/4 {
		*s = append([]interface{}{}, *s...)
	}
	return x
}

// Push item onto the stack.
func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}
