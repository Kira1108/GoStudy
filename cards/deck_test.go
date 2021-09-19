package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
	d := newDeck()
	// test if there are 54 cards in a deck
	if len(d) != 54 {
		t.Errorf("Expected deck length 54, bug got %v", len(d))
	}

	// test if the first card is Heart 1
	if d[0] != "Heart1" {
		t.Errorf("Expected first card Heart1, but got %v", d[0])
	}

	// test if the last card is Joker red
	if d[len(d)-1] != "Joker red" {
		t.Errorf("Expected last card Joker red, but got %v", d[len(d)-1])
	}
}

func TestSaveAndLoad(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(loadedDeck) != 54 {
		t.Errorf("Expected 54 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
