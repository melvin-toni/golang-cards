package main

import (
	"fmt"
	"strings"
)

// Create new type of 'deck' which is a slice of string,
// it's embodiment of card completed with suits and values
type dick []string

// New Deck
func newDeck() dick {
	cards := dick{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver same as 'this' or 'self'
func (d dick) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d dick, handSize int) (dick, dick) {
	return d[:handSize], d[handSize:]
}

func (d dick) toStringYo() string {
	return strings.Join([]string(d), ",")
}