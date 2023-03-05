package main

import (
	"os"
	"testing"
)

// testing newDeck()
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected the len of cards to be 16 but got: ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the Ace of Spades but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected the Four of Hearts but got: %v", d[len(d)-1])
	}
}

// testing saveToFile() and newDeckFromFile()
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	test_file := "_decktesting"
	os.Remove(test_file)

	d := newDeck()

	d.saveToFile(test_file)

	data := newDeckFromFile(test_file)

	if len(data) != 16 {
		t.Errorf("Expected length to be 16 but got: %v", len(data))
	}

	os.Remove(test_file)
}
