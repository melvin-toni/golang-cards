package main

import (
	"os"
	"testing"
)

/*
	test function name we get from function name we want to test
	e.g. if we want to test newDeck() function then test name will be TestNewDeck()
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// if d[0] != "Ace of Spades" {
	// 	t.Errorf("Expected deck of Ace of Spades, but got %v", d[0])
	// }

	// if d[len(d)-1] != "Ace of Spades" {
	// 	t.Errorf("Expected deck of Ace of Spades, but got %v", d[len(d)-1])
	// }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
