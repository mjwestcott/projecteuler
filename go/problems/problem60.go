// problem60.go
//
// https://projecteuler.net/problem=60
//
// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two
// primes and concatenating them in any order the result will always be prime.
// For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of
// these four primes, 792, represents the lowest sum for a set of four primes
// with this property.
//
// Find the lowest sum for a set of five primes for which any two primes
// concatenate to produce another prime.

package main

import (
	"euler/tools"
	"fmt"
	"strconv"
	"sync"
)

type cache struct {
	sync.Mutex
	m map[string]bool
}

func newCache() *cache {
	return &cache{sync.Mutex{}, make(map[string]bool)}
}

func (c *cache) concatsToPrime(x, y int) bool {
	key := strconv.Itoa(x) + strconv.Itoa(y)
	c.Lock()
	defer c.Unlock()
	// Try to find the answer in the cache.
	if val, ok := c.m[key]; ok {
		return val
	}
	// Otherwise find it manually and add to the cache.
	xstr, ystr := strconv.Itoa(x), strconv.Itoa(y)
	a, _ := strconv.Atoi(xstr + ystr)
	b, _ := strconv.Atoi(ystr + xstr)
	val := tools.IsPrime(a) && tools.IsPrime(b)
	c.m[key] = val
	return val
}

func (c *cache) allConcatToPrime(n Node) bool {
	for _, x := range n {
		for _, y := range n {
			if x != y && x < y {
				if c.concatsToPrime(x, y) == false {
					return false
				}
			}
		}
	}
	return true
}

// A Node is a candidate solution to the problem.
type Node []int

// We are going to use a concurrent depth-first search with a worker goroutine
// pool of 4. Each goroutine will search for a solution from a different
// starting prime. As soon as a solution is found, we return from the function.
// Otherwise, we wait for all starting primes to be checked, and return an
// error.
func problem60() (int, error) {
	// It's not clear how many primes to search through. Experimentation
	// suggests a limit of 9000 produces the correct answer: 26033. Note
	// our algorithm does not guarantee the solution is the smallest
	// possible, but as a matter of fact, it is. We could verify our
	// answer by raising the limit to 26033, searching exhaustively, and
	// observing that no smaller solutions are found.
	limit := 9000
	var primes []int
	for i := 0; i < limit; i++ {
		if tools.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	c := newCache()
	ans := make(chan int)   // Used to send the answer.
	done := make(chan bool) // Used to signal that all worker goroutines are done.
	pchan := make(chan int) // Used to send worker goroutines a starting prime to search.
	var wg sync.WaitGroup

	// Woker goroutine pool of 4.
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for {
				p, ok := <-pchan
				if !ok {
					wg.Done()
					return
				}

				// Perform depth-first search starting at the given prime.
				var frontier tools.Stack
				frontier.Push(Node{p})

				for len(frontier) != 0 {
					n := frontier.Pop().(Node)
					if len(n) == 5 {
						ans <- tools.Sum(n...)
					}
					for _, prime := range primes {
						child := append(append(*new(Node), n...), prime)
						if prime > tools.Max(n...) && c.allConcatToPrime(child) {
							frontier.Push(child)
						}
					}
				}
			}
		}()
	}

	go func() {
		for _, p := range primes {
			pchan <- p
		}
		close(pchan)
		wg.Wait()    // Wait for all workers to complete their search
		done <- true // before sending completion signal.
	}()

	for {
		select {
		case x := <-ans:
			return x, nil
		case <-done:
			return -1, fmt.Errorf("problem60: no solution found with limit %v", limit)
		}
	}
}

func main() {
	ans, err := problem60()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
