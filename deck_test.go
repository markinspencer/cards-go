package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as the first card, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected King of Hearts as the last card, but got %v", d[len(d)])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fn := "_decktesting"
	os.Remove(fn)

	d := newDeck()
	d.saveToFile(fn)

	f := newDeckFromFile(fn)

	if len(f) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(f))
	}

	os.Remove(fn)
}
