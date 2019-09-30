/*

	newDeck -> Code to make sure that a deck is created with x number of cards

						-> Create a new deck -> Write if statement to see if the deck has the right number of cards
																	-> If it doesn't, tell the go test handler that something went wrong



					-> Code to make sure that the first card os an Ace of Spades
					-> Code to make sure that the last card is a Four of Clubs

	saveToFile ->
	newDeckFromFile ->

	func TestNewDeck(t *testing.T) {

	}

	t *testing.T ->
		call by the Go testing runner or the go test runner with an argument that we refer to with the name of t and it's a type


		func TestNewDeck(t *testing.T) {
			d := newDeck()

			if len(d) != 52 {
				t.Errorf("Expected deck length of 52, but got %v", len(d))
			}
		}

		Formatted string means that we can pass in some extra values and the how those get automatically injected into the string that we provide as the first argument

		%v -> will be taken from whatever value we passed in right after

*/