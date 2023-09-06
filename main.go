package main

import (
	"cards/deck"
	"fmt"
)

func main() {
	d := deck.FromFile("data.txt")

	// cards.saveToFile("data.txt")
	fmt.Println(d.ToString())
	fmt.Println("After shuffling the deck")
	d.Shuffle()
	fmt.Println(d.ToString())
}
