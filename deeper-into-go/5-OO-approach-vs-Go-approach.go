/*

	Cards - OO Approach

								Deck Instance

	Deck Class -> cards[string]

								print (function)
								shuffle (function)
								saveToFile (function)

		Base Go Types

		String // Integer // Float // Array // Map

		-> We want to "extend" a base type and add some extra functionality to it

			- type deck []string  // Tell Go we want to create an array of strings and add a bunch of functions specifically made to work with it
		->
			- Functions with 'deck' as a 'receiver'	// A function with a receiver is like a "method" - a function that belongs to an "instance"


		Project structure:

			- main.go 			-> Code to create and manipulate a deck
			- deck.go 			-> Code that describes what a deck is and how it works
			- deck_test.go	-> Code to automatically test the deck

*/