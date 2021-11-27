package main

func main() {
	// Go has 2 DS for handling list of records - array(fixed length), slice(dynamic length)
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades") // appends does not moodify existing slice it returns a new one
	cards.print()
}

func newCard() string {
	return "Five of diamond"
}
