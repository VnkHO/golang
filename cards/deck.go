package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
// Receiver function
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

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	sliceString := strings.Split(string(byteSlice), ",")
	return deck(sliceString)
}

func (d deck) shuffle() {

	// Generate a different int64 every single time
	// We use that as the seed to generate a newSource object
	// And then we use that source object as the basis for our new random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) //

	for index := range d {
		newPosition := r.Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
