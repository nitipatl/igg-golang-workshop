package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
