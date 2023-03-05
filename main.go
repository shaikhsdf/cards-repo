package main

func main() {

	// cards := deck{"Ace of spades", "Ace of diamonds"}
	// cards = append(cards, "Ace of hearts")
	cards := newDeck()
	cards.shuffle()

	hand, remCards := deal(cards, 6)
	hand.print()

	remCards.print()

	hand.saveToFile("cards")

	cards_2 := newDeckFromFile("cards")

	cards_2.print()

	cards_2.shuffle()
	cards.print()

	hand_2, remCards_2 := deal(cards_2, 6)
	hand_2.print()
	remCards_2.print()

}
