package main

import "fmt"

// create a new type "deck" which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	cards := d[handSize:]
	return hand, cards
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Heart", "Spade", "Diamond", "Club"}
	cardValue := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
