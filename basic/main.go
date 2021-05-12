package main

import (
	"fmt"

	deck "github.com/nitipatl/go-deck"
	taey_math "github.com/nitipatl/go-math"
)

func main() {
	cards := deck.NewDeck()
	hand, remainingCards := deck.Deal(cards, 5)
	hand.Print()
	remainingCards.Print()
	cards.Print()
	fmt.Println(taey_math.Add(1, 2))
}
