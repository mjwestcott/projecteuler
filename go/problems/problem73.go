// problem73.go
//
// https://projecteuler.net/problem=73
//
// Consider the fraction, n/d, where n and d are positive integers. If n<d and
// HCF(n,d)=1, it is called a reduced proper fraction.
//
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order
// of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 3 fractions between 1/3 and 1/2.
//
// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced
// proper fractions for d ≤ 12,000?

package main

import (
	"euler/tools"
	"fmt"
)

func problem73() int {
	sum := 0
	for d := 1; d <= 12000; d++ {
		for n := 1; n < d; n++ {
			if tools.GCD(n, d) == 1 {
				fn, fd := float64(n), float64(d)
				rat := fn / fd
				if 1.0/3.0 < rat && rat < 0.5 {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(problem73())
}