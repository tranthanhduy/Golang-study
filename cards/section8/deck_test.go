package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	first_d := d[0]
	last_d := d[len(d)-1]

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	fmt.Println(strings.Contains(first_d, "Ace"))

	if !strings.Contains(first_d, "Ace") {
		t.Errorf("First element of Slice is Wrong!")
	}

	fmt.Println(strings.Contains(last_d, "Four"))
	if !strings.Contains(last_d, "Four") {
		t.Errorf("Last element of Slice is Wrong!")
	}
}

func TestSavetoDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	deckload := newDeckFromFile("_decktesting")

	if len(deckload) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deckload))
	}

	os.Remove("_decktesting")
}
