package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// cards := newDeckFromFile("deck.txt")
	// cards.print()
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Intn(10))
	// fmt.Println(time.Now().UnixNano())
}
