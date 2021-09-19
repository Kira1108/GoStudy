package main

import (
	"testing"
)

func TestNewDec(t *testing.T) {
	newDeck := newDeck()
	if len(newDeck) != 54 {
		t.Errorf("Expected deck length 54, bug got %v", len(newDeck))
	}
}
