package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	var TestCards = []Card{
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
	}
	var StartCards = []Card{
		{"7", "Spades", 7},
		{"8", "Spades", 8},
		{"9", "Spades", 9},
		{"10", "Spades", 10},
		{"J", "Spades", 11},
		{"Q", "Spades", 12},
		{"K", "Spades", 13},
		{"A", "Spades", 1},
		{"2", "Spades", 2},
		{"3", "Spades", 3},
		{"4", "Spades", 4},
		{"5", "Spades", 5},
		{"6", "Spades", 6},
	}

	index := func(p1, p2 *Card) bool {
		return p1.index < p2.index
	}
	By(index).Sort(StartCards)

	for i := 0; i < len(TestCards); i++ {
		if StartCards[i] != TestCards[i] {
			t.Errorf("Expected %d, got %d", TestCards[i], StartCards[i])
		}
	}
}
