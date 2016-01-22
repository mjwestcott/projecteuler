// problem62.go
//
// The cube, 41063625 (345**3), can be permuted to produce two other cubes:
// 56623104 (384**3) and 66430125 (405**3). In fact, 41063625 is the smallest cube
// which has exactly three permutations of its digits which are also cube. Find the
// smallest cube for which exactly five permutations of its digits are cube.

package main

import (
	"euler/tools"
	"fmt"
	"math/big"
)

type entry struct {
	count int
	min   *big.Int
}

func problem62() *big.Int {
	cubes := make(map[string]*entry)
	one, three := big.NewInt(1), big.NewInt(3)

	for i := big.NewInt(0); ; i.Add(i, one) {
		x := new(big.Int).Exp(i, three, nil) // x := i**3

		// Sort the digits to produce a 'canonical' representation.
		key := tools.SortedString(x.String())
		e, ok := cubes[key]
		if !ok {
			cubes[key] = &entry{1, x}
		} else {
			e.count++
			if x.Cmp(e.min) == -1 { // x < e.min
				e.min = x
			}
			if e.count == 5 {
				return e.min
			}
		}
	}
}

func main() {
	ans := problem62()
	fmt.Println(ans)
}
