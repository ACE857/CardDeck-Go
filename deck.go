package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func newDeckFromFile(fileName string) deck {
	bsData, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error reading file.", error)
		//return nil
		os.Exit(1)
	}
	return strings.Split(string(bsData), ",")
}

func (d deck) saveDeckToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
