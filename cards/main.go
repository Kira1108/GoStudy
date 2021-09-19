package main

func main() {
	// cards := newDeck()
	cards := readDeckFromFile("cards.txt")
	cards.shuffle()
	cards[:5].print()
	// cards.saveToFile("cards.txt")
	// cards.print()
	// hand, remain := deal(cards, 10)
	// fmt.Println("Cards in hand")
	// hand.print()
	// fmt.Println()
	// fmt.Println("Cards left")
	// remain.print()
}
