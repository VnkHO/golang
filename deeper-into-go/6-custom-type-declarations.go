/*

	Whare we are doing right here is we are saying there is a new thing calles a deck
	It is EQUAL to like all the same behavior as a Slice of String

	class deck extends slices(string)
	type deck []string

	type deck === []string

	// I want to create a new function that belongs to this deck type
	// and whenever we call that function we should print out the each individual card inside of the deck


	// Print -> Purpose of this function is going to be to loop through our deck of cards and print ou the value

	func (d deck) print() {

		for i, card := range d {
			fmt.Println(i, card)
		}
	}

*/