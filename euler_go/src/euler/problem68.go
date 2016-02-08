// problem68.go
//
// https://projecteuler.net/problem=68
//
// What is the maximum 16-digit string for a 'magic' 5-gon ring?

package main

import (
	"euler/tools"
	"fmt"
	"strings"
)

// Ring represents a solution to the 5-gon ring problem.
type Ring []int

// This solution is much less efficient than the one found in the Python
// project dir. There, we created candidate solutions incrementally, allowing
// us to 'exit early' i.e. discount many permutations that are not worth
// considering. It's not quite as convenient to do the same in Go, so we simply
// brute-force check all the factorial(10) permutations.
func problem68() string {
	const start = 6   // Each line cannot sum to less than 6 (1+2+3)
	const end = 27    // or greater than 27 (8+9+10)
	const strlen = 16 // Specified in the problem description.
	var max string
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	perms := tools.Permutations(xs, 10)

	for n := start; n <= end; n++ {
		for _, p := range perms {
			a, b, c, d, e := p[0], p[1], p[2], p[3], p[4]
			f, g, h, i, j := p[5], p[6], p[7], p[8], p[9]
			// Test whether this permutation conforms to the summation rules.
			if (a+b+c != n) || (d+c+e != n) || (f+e+g) != n || (h+g+i) != n || (j+i+b) != n {
				continue
			}
			// Each solution can be described uniquely starting from the group of three
			// with the numerically lowest external node and working clockwise.
			// So we specify that a < min(d, f, h, j)
			if a < tools.Min(d, f, h, j) {
				solution := Ring{
					// As shown in the summation test above, the
					// 5-gon ring is represented as this pattern:
					// {a b c} {d c e} {f e g} {h g i} {j i b}
					// The problem asks for it flattened into one string.
					a, b, c, d, c, e, f, e, g, h, g, i, j, i, b,
				}
				s := strings.Join(tools.IntsToStrings(solution), "")
				if s > max && len(s) == strlen {
					max = s
				}
			}
		}
	}
	return max
}

func main() {
	ans := problem68()
	fmt.Println(ans)
}
