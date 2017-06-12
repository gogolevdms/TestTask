package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Card struct {
	name  string
	suit  string
	index int
}

func (c Card) String() string {
	return fmt.Sprintf("name: %s,\t suit: %s, index: %d\n", c.name, c.suit, c.index)
}

type By func(p1, p2 *Card) bool

func (by By) Sort(cards []Card) {
	ps := &cardSorter{
		cards: cards,
		by:    by,
	}
	sort.Sort(ps)
}

type cardSorter struct {
	cards []Card
	by    func(p1, p2 *Card) bool
}

// Len, Swap and Less is part of sort.Interface.
func (c *cardSorter) Len() int {
	return len(c.cards)
}
func (c *cardSorter) Swap(i, j int) {
	c.cards[i], c.cards[j] = c.cards[j], c.cards[i]
}
func (s *cardSorter) Less(i, j int) bool {
	return s.by(&s.cards[i], &s.cards[j])
}

var cards = []Card{
	{"A", "Spades", 1},
	{"2", "Spades", 2},
	{"3", "Spades", 3},
	{"4", "Spades", 4},
	{"5", "Spades", 5},
	{"6", "Spades", 6},
	{"7", "Spades", 7},
	{"8", "Spades", 8},
	{"9", "Spades", 9},
	{"10", "Spades", 10},
	{"J", "Spades", 11},
	{"Q", "Spades", 12},
	{"K", "Spades", 13},
	{"A", "Clubs", 14},
	{"2", "Clubs", 15},
	{"3", "Clubs", 16},
	{"4", "Clubs", 17},
	{"5", "Clubs", 18},
	{"6", "Clubs", 19},
	{"7", "Clubs", 20},
	{"8", "Clubs", 21},
	{"9", "Clubs", 22},
	{"10", "Clubs", 23},
	{"J", "Clubs", 24},
	{"Q", "Clubs", 25},
	{"K", "Clubs", 26},
	{"A", "Diamonds", 27},
	{"2", "Diamonds", 28},
	{"3", "Diamonds", 29},
	{"4", "Diamonds", 30},
	{"5", "Diamonds", 31},
	{"6", "Diamonds", 32},
	{"7", "Diamonds", 33},
	{"8", "Diamonds", 34},
	{"9", "Diamonds", 35},
	{"10", "Diamonds", 36},
	{"J", "Diamonds", 37},
	{"Q", "Diamonds", 38},
	{"K", "Diamonds", 39},
	{"A", "Hearts", 40},
	{"2", "Hearts", 41},
	{"3", "Hearts", 42},
	{"4", "Hearts", 43},
	{"5", "Hearts", 44},
	{"6", "Hearts", 45},
	{"7", "Hearts", 46},
	{"8", "Hearts", 47},
	{"9", "Hearts", 48},
	{"10", "Hearts", 49},
	{"J", "Hearts", 50},
	{"Q", "Hearts", 51},
	{"K", "Hearts", 52},
}

func main() {
	rand.Seed(100)
	for i := 0; i < 52; i++ {
		j := rand.Intn(len(cards))
		cards[i], cards[j] = cards[j], cards[i]
	}
	fmt.Println(cards)

	index := func(p1, p2 *Card) bool {
		return p1.index < p2.index
	}
	By(index).Sort(cards)
	fmt.Println(cards)
}
