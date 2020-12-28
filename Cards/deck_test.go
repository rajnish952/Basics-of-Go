package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // Go automatically calls test with (t *testing.T) parameter -> t is test handler
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}
}

func TestSavetoFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck length of 20 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
