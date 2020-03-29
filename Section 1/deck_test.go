package Section_1

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting")


	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck :=newDeckFromFile("_decktesting")

	if len(loadedDeck) != 62 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}


	os.Remove("_decktesting")
}