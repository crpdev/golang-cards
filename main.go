package main

var deckSize int

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println()
	// hand, remainingDeck := cards.deal(5)
	// hand.print()
	// fmt.Println()
	// remainingDeck.print()
	// fmt.Println()
	// fmt.Println(cards.asString())
	// cards.saveToFile("cards.txt")
	// cards = newDeckFromFile("card.txt")
	// fmt.Println("From File")
	// cards.print()
}
