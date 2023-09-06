package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func newDeck() deck {
	cards := deck{}

	modifiers := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, modifier := range modifiers {
		for _, value := range values {
			cards = append(cards, value+" of "+modifier)
		}
	}

	return cards
}
