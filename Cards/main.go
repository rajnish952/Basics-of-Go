package main

func main() {

	cards := newDeck()

	// Return two deck types, one assigned to hands other to remaining
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	// fmt.Println(hand.toString())
	cards.saveToFile("MyCards.txt")
	newCards := newDeckFromFile("MyCards.txt")
	newCards.shuffle()
	newCards.print()
	// fmt.Println(newCards)
}
