package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Deck length invalid expected 52 but got $%v", len(d))
	}
	if d[0] != "Ace of Heart" {
		t.Errorf("Deck card[0] invalid expected 'Ace of Heart' but got $%v", d[0])
	}
	if d[len(d)-1] != "Ten of Club" {
		t.Errorf("Deck card[51] invalid expected 'Ten of Club' but got $%v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	d := newDeck()
	d.saveDeckToFile("_decktesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(d) != len(loadedDeck) {
		t.Errorf("Error deck length does not match")
	}
	os.Remove("_decktesting")
}
