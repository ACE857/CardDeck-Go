package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Deck length invalid expected 52 but got $%v", len(d))
	}
}
