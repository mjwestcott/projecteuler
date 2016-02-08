// problem61.go
//
// https://projecteuler.net/problem=61
//
// Triangle, square, pentagonal, hexagonal, heptagonal, and octagonal numbers are
// all figurate (polygonal) numbers and are generated by the following formulae:
//
//     Triangle   P[3,n]=n(n+1)/2  --> 1, 3, 6, 10, 15, ...
//     Square     P[4,n]=n2        --> 1, 4, 9, 16, 25, ...
//     Pentagonal P[5,n]=n(3n−1)/2 --> 1, 5, 12, 22, 35, ...
//     Hexagonal  P[6,n]=n(2n−1)   --> 1, 6, 15, 28, 45, ...
//     Heptagonal P[7,n]=n(5n−3)/2 --> 1, 7, 18, 34, 55, ...
//     Octagonal  P[8,n]=n(3n−2)   --> 1, 8, 21, 40, 65, ...
//
// The ordered set of three 4-digit numbers: 8128, 2882, 8281, has
// three interesting properties.
//
//     1. The set is cyclic, in that the last two digits of each number is the
//     first two digits of the next number (including the last number with the
//     first).
//     2. Each polygonal type: triangle (P[3,127]=8128), square (P[4,91]=8281),
//     and pentagonal (P[5,44]=2882), is represented by a different number in the
//     set.
//     3. This is the only set of 4-digit numbers with this property.
//
// Find the sum of the only ordered set of six cyclic 4-digit numbers for which
// each polygonal type: triangle, square, pentagonal, hexagonal, heptagonal, and
// octagonal, is represented by a different number in the set.

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
)

func triangle(n int) int   { return n * (n + 1) / 2 }
func square(n int) int     { return n * n }
func pentagonal(n int) int { return n * (3*n - 1) / 2 }
func hexagonal(n int) int  { return n * (2*n - 1) }
func heptagonal(n int) int { return n * (5*n - 3) / 2 }
func octagonal(n int) int  { return n * (3*n - 2) }

// Are these four-digits numbers cyclic?
func isCyclic(x, y int) bool {
	return strconv.Itoa(x)[2:] == strconv.Itoa(y)[:2]
}

// Is the set of digits cyclic, including the last number with the first?
func isAllCyclic(n Node) bool {
	z := len(n)
	for i := range n {
		if !isCyclic(n[i], n[(i+1)%z]) {
			return false
		}
	}
	return true
}

// Given a function representing a polygonal type, return all four-digit
// members of that type.
func fourDigitPolys(f func(x int) int) []int {
	var res []int
	for i := 0; ; i++ {
		x := f(i)
		z := len(strconv.Itoa(x))
		switch {
		case z == 4:
			res = append(res, x)
		case z > 4:
			return res
		}
	}
}

func allFourDigitPolygons() [][]int {
	var res [][]int
	funcs := []func(int) int{
		triangle, square, pentagonal,
		hexagonal, heptagonal, octagonal,
	}
	for _, f := range funcs {
		polys := fourDigitPolys(f)
		res = append(res, polys)
	}
	return res
}

// A Node is a candidate solution to the problem.
type Node []int

// For every way to order the polygonal types, perform depth-first search,
// returning as soon as a solution is found.
func problem61() int {
	polys := allFourDigitPolygons()
	perms := tools.Permutations([]int{0, 1, 2, 3, 4, 5}, 6)
	for _, perm := range perms {
		// We will build up candidate solutions incrementally. The
		// starting nodes are therefore the members of polys[perm[0]].
		var frontier tools.Stack
		for _, p := range polys[perm[0]] {
			frontier.Push(Node{p})
		}
		for len(frontier) > 0 {
			n := frontier.Pop().(Node)
			z := len(n)
			if z == 6 && isAllCyclic(n) {
				return tools.Sum(n...)
			} else if z < 6 {
				// Note that perm[z] contains the next
				// polygonal type in this order.
				for _, y := range polys[perm[z]] {
					if isCyclic(n[z-1], y) {
						child := append(append(Node{}, n...), y)
						frontier.Push(child)
					}
				}
			}
		}
	}
	return -1 // Unreachable because there is a known solution.
}

func main() {
	ans := problem61()
	fmt.Println(ans)
}
