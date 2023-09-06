package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	str := strings.Join(d, "\n")

	return str
}

func (d deck) toByteSlice() []byte {
	str := d.toString()

	return []byte(str)
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlice(), 0666)
}

func fromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		// option #1: print the error message and force app to quit with exit code 1
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	str := strings.Split(string(bytes), "\n")

	return str
}
