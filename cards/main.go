package main

func main() {
	cards := newDeck()
	// deal(cards, 3)
	// cards.toString()
	// cards.saveToFile("cards.txt")
	// cardsFromFile := newDeckFromFile("cards.txt")
	// cardsFromFile.print()

	cards.shuffle()
	cards.print()
}
