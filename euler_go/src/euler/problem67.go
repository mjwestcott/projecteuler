// problem67.go
//
// Find the maximum total from top to bottom in triangle.txt a 15K text
// file containing a triangle with one-hundred rows.

package main

import (
	"bufio"
	"euler/tools"
	"fmt"
	"log"
	"os"
	"strings"
)

// Triangle is a list of lists of ints such as [[1] [2 3] [4 5 6]]
// representing a triangle of the form:
//     1
//    2 3
//   4 5 6
type Triangle [][]int

var cache = make(map[string]int)

// Return a key suitable for use in a map.
// (Note: currently very slow.)
func k(tr Triangle) string {
	return fmt.Sprintf("%v", tr)
}

// Recursively find the maximum value of the root node plus the largest
// of its children, and so on, all the way to the base.
func maxRoute(tr Triangle) int {
	if val, ok := cache[k(tr)]; ok {
		return val
	}
	root := tr[0][0]
	if len(tr) == 1 {
		return root
	}
	a, b := children(tr)
	result := root + tools.Max(maxRoute(a), maxRoute(b))
	cache[k(tr)] = result
	return result
}

// Split the triangle in two below the root node.
// e.g. [[1] [2 3] [4 5 6]] -> [[2] [4 5]] [[3] [5 6]]
func children(tr Triangle) (Triangle, Triangle) {
	var a, b Triangle
	for _, row := range tr[1:] {
		a = append(a, row[:len(row)-1])
		b = append(b, row[1:])
	}
	return a, b
}

func problem67() int {
	file, err := os.Open("data/triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tr Triangle
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		ss := strings.Split(row, " ")
		tr = append(tr, tools.StringsToInts(ss))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return maxRoute(tr)
}

func main() {
	ans := problem67()
	fmt.Println(ans)
}
