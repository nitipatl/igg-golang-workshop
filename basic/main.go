package main

func main() {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	cards.print()
}
