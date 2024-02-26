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

	if d[0] != "ace of spades" {
		t.Errorf("Expected first card of deck to be ace of spades, but got %v", d[0])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	fileName := "_decktesting.txt"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
