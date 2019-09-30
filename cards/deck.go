package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

/*
	Whare we are doing right here is we are saying there is a new thing calles a deck
	It is EQUAL to like all the same behavior as a Slice of String
*/
// class deck extends slices(string)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print -> Purpose of this function is going to be to loop through our deck of cards and print ou the value
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
