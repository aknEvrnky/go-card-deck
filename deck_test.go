package main

import (
	"cards/deck"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()

	if len(d) != 13*4 {
		t.Errorf("Expected number of 52 cards in a deck. %v got.", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be \"Ace of Spaces\". Got \"%s\"", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected the last card to be \"King of Clubs\". Got \"%s\"", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	tempFile := t.TempDir() + "/_decktesting.txt"

	d := deck.NewDeck()

	d.SaveToFile(tempFile)

	loadedDeck := deck.FromFile(tempFile)

	if len(d) != len(loadedDeck) {
		t.Error("number of items of saved deck and loaded deck did not match.")
	}
}
