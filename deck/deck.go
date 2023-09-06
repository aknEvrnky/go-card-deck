package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Deck []string

func (d Deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func NewDeck() Deck {
	cards := Deck{}

	modifiers := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, modifier := range modifiers {
		for _, value := range values {
			cards = append(cards, value+" of "+modifier)
		}
	}

	return cards
}

func (d Deck) ToString() string {
	str := strings.Join(d, "\n")

	return str
}

func (d Deck) ToByteSlice() []byte {
	str := d.ToString()

	return []byte(str)
}

func (d Deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, d.ToByteSlice(), 0666)
}

func FromFile(filename string) Deck {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		// option #1: print the error message and force app to quit with exit code 1
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	str := strings.Split(string(bytes), "\n")

	return str
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
