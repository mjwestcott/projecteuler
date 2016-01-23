// problem64.go
//
// The first ten continued fraction representations of (irrational) square roots are:
//
//     sqrt(2)=[1;(2)]          period=1
//     sqrt(3)=[1;(1,2)]        period=2
//     sqrt(5)=[2;(4)]          period=1
//     sqrt(6)=[2;(2,4)]        period=2
//     sqrt(7)=[2;(1,1,1,4)]    period=4
//     sqrt(8)=[2;(1,4)]        period=2
//     sqrt(10)=[3;(6)]         period=1
//     sqrt(11)=[3;(3,6)]       period=2
//     sqrt(12)=[3;(2,6)]       period=2
//     sqrt(13)=[3;(1,1,1,1,6)] period=5
//
// Exactly four continued fractions, for N <= 13, have an odd period. How many
// continued fractions for N <= 10000 have an odd period?

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func continuedFracSqrt(S float64) []int {
	// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
	// Using variables S, m, d, a as in the URL above.
	m := 0.0
	d := 1.0
	a := math.Floor(math.Sqrt(S))
	result := []int{}
	seen := make(map[string]bool)
	for {
		seen[k([]float64{m, d, a})] = true
		result = append(result, int(a))

		m = (d * a) - m
		d = (S - math.Pow(m, 2.0)) / d
		if d == 0 {
			// S is a perfect square.
			return result
		}
		a = math.Floor((math.Floor(math.Sqrt(S)) + m) / d)

		// The algorithm terminates when [m d a] repeats.
		if seen[k([]float64{m, d, a})] {
			return result
		}
	}

}

// Return a key suitable for use in a map.
func k(xs []float64) string {
	var ss []string
	for _, x := range xs {
		ss = append(ss, strconv.FormatFloat(x, 'f', -1, 64))
	}
	return strings.Join(ss, " ")
}

func problem64() int {
	count := 0
	for i := 2; i <= 10000; i++ {
		// The first element is not part of the period.
		period := continuedFracSqrt(float64(i))[1:]
		if len(period)%2 == 1 {
			count++
		}
	}
	return count
}

func main() {
	ans := problem64()
	fmt.Println(ans)
}
