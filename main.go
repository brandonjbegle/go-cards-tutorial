package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("mycards.txt")
	//fmt.Println(cards)
}
