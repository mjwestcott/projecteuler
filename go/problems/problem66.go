// problem66.go
//
// https://projecteuler.net/problem=66
//
// Consider quadratic Diophantine equations of the form: x**2 – Dy**2 = 1.
// For example, when D=13, the minimal solution in x is 6492 – 13×1802 = 1.
// It can be assumed that there are no solutions in positive integers when
// D is square. By finding minimal solutions in x for D = {2, 3, 5, 6, 7},
// we obtain the following:
//
//     3**2 – 2×2**2 = 1
//     2**2 – 3×1**2 = 1
//     9**2 – 5×4**2 = 1
//     5**2 – 6×2**2 = 1
//     8**2 – 7×3**2 = 1
//
// Hence, by considering minimal solutions in x for D ≤ 7, the largest x is
// obtained when D=5.  Find the value of D ≤ 1000 in minimal solutions of x
// for which the largest value of x is obtained.

package main

import (
	"fmt"
	"math"
	"math/big"
)

// Each iteration through the convergents of the continued fraction of sqrt(D),
// we want to check whether the numerator and denominator provide a solution to
// the Diophantine equation: https://en.wikipedia.org/wiki/Pell%27s_equation
// See the section entitled 'Fundamental solution via continued fractions'

// Return n values in the continued fraction representation of sqrt(D),
// e.g. sqrt(23) = [4;(1,3,1,8)] thus processCF(23, 10) == [4 1 3 1 8 1 3 1 8 1]
// See problem64.go for a link explaining the algorithm.
func processCF(D, n int) []int {
	var result []int
	S := float64(D)
	m, d, a := 0.0, 1.0, math.Floor(math.Sqrt(float64(S)))

	for ; n > 0; n-- {
		result = append(result, int(a))
		m = (d * a) - m
		d = (S - math.Pow(m, 2.0)) / d
		a = math.Floor((math.Floor(math.Sqrt(S)) + m) / d)
	}
	return result
}

// Return the nth convergent of the continued fraction for sqrt(D),
// where D is a non-square positive integer.
func convergent(D, n int) *big.Rat {
	if n == 1 {
		first := processCF(D, 1)[0]
		return big.NewRat(int64(first), 1)
	}
	// Collect the first n partial values of D.
	values := processCF(D, n)
	// Construct the continued fraction, where 'tail' is the recursive component.
	next, values := values[0], values[1:]
	x := big.NewRat(int64(next), 1)
	one := big.NewRat(1, 1)
	return new(big.Rat).Add(x, new(big.Rat).Quo(one, tail(values)))
}

// Recursively return the tail end of the continued fraction for sqrt(D).
func tail(values []int) *big.Rat {
	next, values := values[0], values[1:]
	x := big.NewRat(int64(next), 1)
	if len(values) == 0 {
		return x
	}
	one := big.NewRat(1, 1)
	return new(big.Rat).Add(x, new(big.Rat).Quo(one, tail(values)))
}

// Find the solution with the minimal value of x satisfying the equation.
func solvePellsEquation(D int) *big.Int {
	for i := 1; ; i++ {
		candidate := convergent(D, i)
		if isSolution(D, candidate) {
			// We only need the value of x for this problem.
			return candidate.Num()
		}
	}
}

// Do the numerator and denominator of r provide a solution to the equation
// x**2 - D*(y**2) = 1?
func isSolution(D int, r *big.Rat) bool {
	S := big.NewInt(int64(D))
	x := r.Num()
	y := r.Denom()

	one := big.NewInt(1)
	two := big.NewInt(2)

	a := new(big.Int).Exp(x, two, nil)
	b := new(big.Int).Exp(y, two, nil)
	return new(big.Int).Sub(a, new(big.Int).Mul(S, b)).Cmp(one) == 0
}

func problem66() int {
	max := new(big.Int)
	D := 0

	for i := 1; i <= 1000; i++ {
		z := math.Sqrt(float64(i))
		if math.Abs(z-math.Floor(z)) == 0 {
			// i is a perfect square, ignore it.
			continue
		}
		x := solvePellsEquation(i)
		// Keep track of the largest x value seen, and
		// the value of D for which it was obtained.
		if x.Cmp(max) == 1 {
			max, D = x, i
		}
	}
	return D
}

func main() {
	fmt.Println(problem66())
}
