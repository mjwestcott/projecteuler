// problem60.go
//
// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two
// primes and concatenating them in any order the result will always be prime.
// For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of
// these four primes, 792, represents the lowest sum for a set of four primes
// with this property.
//
// Find the lowest sum for a set of five primes for which any two primes
// concatenate to produce another prime.

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
)

// Used to memoize concatsToPrime.
var cache = make(map[string]bool)

// Translate the args to a string key to be used in the cache.
func k(x, y int) string {
	return strconv.Itoa(x) + strconv.Itoa(y)
}

func concatsToPrime(x, y int) bool {
	if v, ok := cache[k(x, y)]; ok {
		return v
	}
	xstr, ystr := strconv.Itoa(x), strconv.Itoa(y)
	a, _ := strconv.Atoi(xstr + ystr)
	b, _ := strconv.Atoi(ystr + xstr)
	v := tools.IsPrime(a) && tools.IsPrime(b)
	cache[k(x, y)] = v
	return v
}

func allConcatToPrime(candidates []int) bool {
	for _, x := range candidates {
		for _, y := range candidates {
			if x != y && x < y {
				if concatsToPrime(x, y) == false {
					return false
				}
			}
		}
	}
	return true
}

// A node is a candidate solution to the problem.
type node []int

func (n node) max() int {
	if len(n) == 0 {
		panic("max: node has 0 elements")
	}
	x := n[0]
	for _, y := range n[1:] {
		if y > x {
			x = y
		}
	}
	return x
}

func (n node) sum() int {
	res := 0
	for _, x := range n {
		res += x
	}
	return res
}

// Stack is a simple stack implementation.
type stack []node

// Pop an item off the stack. Panics if s is empty.
func (s *stack) pop() node {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	// Shrink the underlying array if the slice len <= 1/4 its size.
	if len(*s) <= cap(*s)/4 {
		*s = append([]node{}, *s...)
	}
	return x
}

// Push an item onto the stack.
func (s *stack) push(x node) {
	*s = append(*s, x)
}

func problem60() (int, error) {
	// It's not clear how many primes to search through. Experimentation
	// suggests 9000.
	limit := 9000
	var primes []int
	for i := 0; i < limit; i++ {
		if tools.IsPrime(i) {
			primes = append(primes, i)
		}
	}
	tools.ReverseInts(primes) // We want to search smaller primes first from pop().

	// Use depth-first search.
	var frontier stack
	for _, p := range primes {
		frontier.push(node{p})
	}
	for len(frontier) != 0 {
		n := frontier.pop()
		if len(n) == 5 {
			return n.sum(), nil
		}
		for _, p := range primes {
			child := append(append(*new(node), n...), p)
			if p > n.max() && allConcatToPrime(child) {
				frontier.push(child)
			}
		}
	}
	return -1, fmt.Errorf("problem60: no solution found, limit=%d", limit)
}

func main() {
	ans, _ := problem60()
	fmt.Println(ans)
}
