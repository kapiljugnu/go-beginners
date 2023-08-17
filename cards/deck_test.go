package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected last card of Four of clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)

	lodedDeck := newDeckFromFile(filename)
	if len(lodedDeck) != len(d) {
		t.Errorf("Expected %v cards in deck, got %v", len(d), len(lodedDeck))
	}

	os.Remove(filename)
}
