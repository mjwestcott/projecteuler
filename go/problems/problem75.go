// problem75.go
//
// https://projecteuler.net/problem=75
//
// It turns out that 12 cm is the smallest length of wire that can be bent to form
// an integer sided right angle triangle in exactly one way, but there are many
// more examples.
//
//     12 cm: (3,4,5)
//     24 cm: (6,8,10)
//     30 cm: (5,12,13)
//     36 cm: (9,12,15)
//     40 cm: (8,15,17)
//     48 cm: (12,16,20)
//
// In contrast, some lengths of wire, like 20 cm, cannot be bent to form an integer
// sided right angle triangle, and other lengths allow more than one solution to be
// found; for example, using 120 cm it is possible to form exactly three different
// integer sided right angle triangles.
//
//     120 cm: (30,40,50), (20,48,52), (24,45,51)
//
// Given that L is the length of the wire, for how many values of L ≤ 1,500,000 can
// exactly one integer sided right angle triangle be formed?

package main

import (
	"euler/tools"
	"fmt"
)

// Triple represents a Pythagorean triple e.g. (3, 4, 5)
type Triple []int

// Given a Pythagorean triple, return its three children triples.
func (t Triple) children() []Triple {
	// See Berggren's ternary tree, which will produce all infinitely many
	// primitive triples without dupication.
	a, b, c := t[0], t[1], t[2]
	a1, b1, c1 := (-a + 2*b + 2*c), (-2*a + b + 2*c), (-2*a + 2*b + 3*c)
	a2, b2, c2 := (+a + 2*b + 2*c), (+2*a + b + 2*c), (+2*a + 2*b + 3*c)
	a3, b3, c3 := (+a - 2*b + 2*c), (+2*a - b + 2*c), (+2*a - 2*b + 3*c)
	return []Triple{{a1, b1, c1}, {a2, b2, c2}, {a3, b3, c3}}
}

func problem75() int {
	limit := 1500000

	// A mapping from values of L to the number of right-angles triangles
	// with the perimiter L.
	triangles := make(map[int]int)

	// Use a depth-first search to exhaust the search space, starting with
	// the first Pythagorean triple.
	var frontier tools.Stack
	frontier.Push(Triple{3, 4, 5})

	for len(frontier) > 0 {
		t := frontier.Pop().(Triple)
		L := tools.Sum(t...)
		if L > limit {
			continue
		}
		triangles[L]++

		// We're not only interested in 'primitive triples', but
		// multiples too
		a, b, c := t[0], t[1], t[2]
		for i := 2; ; i++ {
			multiple := Triple{i * a, i * b, i * c}
			if tools.Sum(multiple...) >= limit {
				break
			}
			triangles[tools.Sum(multiple...)]++
		}

		for _, child := range t.children() {
			frontier.Push(child)
		}
	}

	count := 0
	for _, val := range triangles {
		if val == 1 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(problem75())
}
