/*

	Create a deck
	-> Save to file
	-> File created!
	-> Attempt to lad file
	-> Crash!

	We created the file at File created!
	And so if we crash at Crash! we never really get the opportunity to do any cleanup or remove
	that temporary testing file that we had just created to test saving a file to the hard ddrive

	So we always need to make sure that whenever we are writing thest with Go we take care of any cleanup that needs to be done very manually
	because there is absolutely no hand-holding by the Go testing framework

	The Go testing framework is not going to detect that you wrote some file and automatically attempt to clean it up for you

	You always have to take care of clean up yourself

	Testing saveToDeck and newDeckFromFile

		-> Delete any files in current working directory with the name "_decktesting"
		-> Create a deck
		-> Save to file "_decktesting"
		-> Load from file
		-> Assert deck length
		-> Delete any files in current working directory with the name "_decktesting"


	func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
		// Remove file
		os.Remove("_decktesting")

		d := newDeck()
		d.saveToFile("_decktesting")

		loadedDeck := newDeckFromFile("_decktesting")

		if len(loadedDeck) != 52 {
			t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
		}

		os.Remove("_decktesting")
	}

*/