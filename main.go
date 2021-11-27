package main

import "fmt"

func main() {
	//var card = "Ace of spades"
	card := "Ace of spades" // := is only used when defining new var, when reassigning only use =
	// variables can be initilized outside the function be can only be assigned insde
	card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of diamond"
}
