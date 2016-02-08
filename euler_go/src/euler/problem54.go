// problem54.go
//
// https://projecteuler.net/problem=54
//
// The file, poker.txt, contains one-thousand random hands dealt to two
// players. How many hands does Player 1 win?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Hand represents a five-card poker hand.
type Hand struct {
	// The ranks are sorted in descending order.
	// e.g. the hand [AC 8D 8H 3S 2S] has ranks [14 8 8 3 2]
	//               [KC 9D 9H 3C 2C]           [13 9 9 3 2]
	ranks []int

	// A set of strings which the hand contains. Used to determine if
	// the hand is a flush.
	suits map[string]bool

	// A slice of [count rank] pairs. This will be used to break ties between
	// two hands of equal value (e.g. two hands that both have 'one-pair').
	// Will be sorted highest count first, then highest rank first.
	// e.g. the hands [AC 8D 8H 3S 2S] and [KC 9D 9H 3C 2C] have groups
	//	[[2 8], [1 14], [1 3], [1 2]]
	//	[[2 9], [1 13], [1 3], [1 2]]
	// This allows us to correctly judge that the second hand beats the former
	// by sorting Hands lexicographically based on the groups field.
	groups [][]int

	// A sorted slice of counts of the card ranks.
	// e.g. the ranks [14 8 8 3 2] have pattern [2 1 1 1]
	//                [10 9 7 7 7]              [3 1 1]
	// Used to determine whether the hand has 'two-pair' or 'threeofakind' etc.
	pattern []int
}

// newHand constructs a new Hand type from the given cards.
// cards is of the form e.g. [AC 8D 8H 3C 2S]
func newHand(cards []string) Hand {
	trans := map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10}

	var ranks []int                  // The ranks field.
	suits := make(map[string]bool)   // The suits field.
	rankToCount := make(map[int]int) // Used to create other fields.
	for _, card := range cards {
		s := string(card[1])
		suits[s] = true
		r, ok := trans[string(card[0])]
		if !ok {
			r, _ = strconv.Atoi(string(card[0]))
		}
		ranks = append(ranks, r)
		rankToCount[r]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))

	// An ace should be played 'low' i.e. with rank 1 if it makes a straight.
	if reflect.DeepEqual(ranks, []int{14, 5, 4, 3, 2}) {
		ranks = []int{5, 4, 3, 2, 1}
	}

	var pattern []int                   // The pattern field.
	countToRanks := make(map[int][]int) // Used to create the groups field.
	for k, v := range rankToCount {
		pattern = append(pattern, v)
		countToRanks[v] = append(countToRanks[v], k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pattern)))

	var groups [][]int         // The groups field.
	seen := make(map[int]bool) // Used to ensure no duplicates.
	for _, count := range pattern {
		// For every count in the pattern append to groups a [count rank] pair.
		// groups is ordered by highest count first, then highest rank first.
		if !seen[count] {
			rs := countToRanks[count]
			sort.Sort(sort.Reverse(sort.IntSlice(rs)))
			for _, r := range rs {
				groups = append(groups, []int{count, r})
			}
		}
		seen[count] = true
	}

	return Hand{ranks, suits, groups, pattern}
}

// Hand type predicates.
func (h Hand) onepair() bool       { return isEqual(h.pattern, []int{2, 1, 1, 1}) }
func (h Hand) twopair() bool       { return isEqual(h.pattern, []int{2, 2, 1}) }
func (h Hand) threeofakind() bool  { return isEqual(h.pattern, []int{3, 1, 1}) }
func (h Hand) fourofakind() bool   { return isEqual(h.pattern, []int{4, 1}) }
func (h Hand) fullhouse() bool     { return isEqual(h.pattern, []int{3, 2}) }
func (h Hand) flush() bool         { return len(h.suits) == 1 }
func (h Hand) straight() bool      { return (len(h.pattern) == 5) && (h.ranks[0]-h.ranks[4] == 4) }
func (h Hand) straightflush() bool { return h.flush() && h.straight() }

// Check whether two []int are identical.
func isEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Return a value for hand from an eight-point scale.
func (h Hand) evaluate() int {
	switch {
	case h.straightflush():
		return 8
	case h.fourofakind():
		return 7
	case h.fullhouse():
		return 6
	case h.flush():
		return 5
	case h.straight():
		return 4
	case h.threeofakind():
		return 3
	case h.twopair():
		return 2
	case h.onepair():
		return 1
	default:
		return 0
	}
}

// Did Player 1 win this round?
func player1wins(h1, h2 Hand) bool {
	// First check whether hands have different value on eight-point scale.
	if v1, v2 := h1.evaluate(), h2.evaluate(); v1 != v2 {
		return v1 > v2
	}
	// If those values are equal, perform lexicographic comparison based on
	// the groups field, a slice of [count rank] pairs ordered by highest
	// count first, then highest rank first.
	for i := range h1.groups {
		// Since v1 == v2, the groups are the same length and the
		// counts are identical. Therefore, compare ranks.
		g1, g2 := h1.groups[i], h2.groups[i]
		rank1, rank2 := g1[1], g2[1]
		switch {
		case rank1 > rank2:
			return true
		case rank1 < rank2:
			return false
		}
	}
	// The problem specifies that ties are not possible, so this
	// should be unreachable. This code path corresponds to
	// two hands being equal in all relevant respects (i.e. they
	// only differ by suit).
	return false
}

func problem54() int {
	file, err := os.Open("data/poker.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int // Count of number of times Player 1 wins
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		cards := strings.Split(row, " ")
		h1, h2 := newHand(cards[:5]), newHand(cards[5:])
		if player1wins(h1, h2) {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

func main() {
	ans := problem54()
	fmt.Println(ans)
}
