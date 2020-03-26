package main

import "testing"

// all tests have to have _test at the end the name
// "go test" will run the tests
// *testing.T
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Deck is impropper size: Expected 52, Got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Deck is out of order. Expected first card of Ace of Spades, got: %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Deck is out of order. Expected last card of King of Clubs, got: %v", d[len(d)-1])
	}
}
