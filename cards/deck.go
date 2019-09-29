package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

/*
	Whare we are doing right here is we are saying there is a new thing calles a deck
	It is EQUAL to like all the same behavior as a Slice of String
*/
// class deck extends slices(string)

type deck []string

// Print -> Purpose of this function is going to be to loop through our deck of cards and print ou the value

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
