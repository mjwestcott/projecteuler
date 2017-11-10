// problem68.go
//
// https://projecteuler.net/problem=68
//
// What is the maximum 16-digit string for a 'magic' 5-gon ring?

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
	"strings"
)

// A brute force method would generate all 10! permutations of the integers
// 1 through 10 and filter for correct solutions. It would be better to
// generate the candidate solutions incrementally so that we ignore as many
// irrelevant permutations as possible.
//
// If we name the ten unique integers of a candidate solution a to j, then the
// problem specifies we must arrange them into five three-element lists as
// follows:
//     [a b c] [d c e] [f e g] [h g i] [j i b]
// such that each list sums to some given number N. It is convenient for us
// to represent the candidate solution as a flat 15 element slice of ints:
//     [a b c d c e f e g h g i j i b]

// Ring represents a candidate solution
type Ring struct {
	solution []int        // the candidate solution being created
	used     map[int]bool // set of ints used in our solution
	N        int          // the sum constraint
}

// Given an incomplete Ring of length l, return all possible length l+1 Rings
func successors(r Ring) []Ring {
	var rs []Ring
	l := len(r.solution)
	switch l {
	case 0, 1, 3, 6, 9, 12:
		// The only constraint on these elements is that they dont
		// match any other elements.
		for i := 1; i <= 10; i++ {
			if !r.used[i] {
				rs = append(rs, makeChild(r, i))
			}
		}
	case 2, 5, 8, 11:
		// These elements must also satisfy the constraint that its three-element
		// group sums to N.
		for i := 1; i <= 10; i++ {
			if !r.used[i] && (i+r.solution[l-1]+r.solution[l-2]) == r.N {
				rs = append(rs, makeChild(r, i))
			}
		}
	case 4, 7, 10, 13:
		// The elements are repeated.
		rs = append(rs, makeChild(r, r.solution[l-2]))
	case 14:
		// This element is a special case repetition.
		rs = append(rs, makeChild(r, r.solution[1]))
	}
	return rs
}

// Create a new Ring object from the parent r, adding the int i to the
// candidate solution.
func makeChild(r Ring, i int) Ring {
	child := Ring{
		solution: make([]int, len(r.solution)),
		used:     make(map[int]bool),
		N:        r.N,
	}
	for k, v := range r.used {
		child.used[k] = v
	}
	child.used[i] = true
	copy(child.solution, r.solution)
	child.solution = append(child.solution, i)
	return child
}

func problem68() int {
	const start = 6 // Each line cannot sum to less than 6 (1+2+3)
	const end = 27  // or greater than 27 (8+9+10)

	// Use a depth-first search.
	var frontier tools.Stack
	var solutions [][]int

	// Populate the initial frontier.
	for N := start; N <= end; N++ {
		for i := 1; i <= 10; i++ {
			r := Ring{
				solution: []int{i},
				used:     make(map[int]bool),
				N:        N,
			}
			r.used[i] = true
			frontier.Push(r)
		}
	}

	for len(frontier) > 0 {
		r := frontier.Pop().(Ring)
		if len(r.solution) == 15 {
			// Each solution can be described uniquely starting from the
			// group of three with the numerically lowest external node and
			// working clockwise. So we need to filter for solutions where
			// a < min(d, f, h, j)
			s := r.solution
			if s[0] < tools.Min(s[3], s[6], s[9], s[12]) {
				solutions = append(solutions, s)
			}
		}
		for _, child := range successors(r) {
			frontier.Push(child)
		}
	}

	// Process solutions according to the problem spec.
	var strlen = 16
	var answers []int
	for _, sol := range solutions {
		var ss []string
		for _, x := range sol {
			ss = append(ss, strconv.Itoa(x))
		}
		s := strings.Join(ss, "")
		if len(s) == strlen {
			val, _ := strconv.Atoi(s)
			answers = append(answers, val)
		}
	}
	return tools.Max(answers...)
}

func main() {
	fmt.Println(problem68())
}
