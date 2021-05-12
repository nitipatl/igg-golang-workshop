package main

import (
	"./deck"
)

func main() {
	cards := deck.NewDeck()
	hand, remainingCards := deck.Deal(cards, 5)
	hand.Print()
	remainingCards.Print()
	cards.Print()
}
