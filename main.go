package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("Dealing hand of size 5")
	hand, cards := deal(cards, 5)
	fmt.Println(hand, cards)
}
