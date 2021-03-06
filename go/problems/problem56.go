// problem56.go
//
// https://projecteuler.net/problem=56
//
// Considering natural numbers of the form, a**b, where a, b < 100,
// what is the maximum digital sum?

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func digitSum(n *big.Int) int64 {
	var sum int64
	for _, r := range n.String() {
		x, _ := strconv.Atoi(string(r))
		sum += int64(x)
	}
	return sum
}

func problem56() int64 {
	var max, i, j int64
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			a, b, c := big.NewInt(i), big.NewInt(j), new(big.Int)
			c.Exp(a, b, nil) // c = a**b
			sum := digitSum(c)
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	fmt.Println(problem56())
}
