/*

	for each index, cards in cards
		Generate a random number between 0 and len(cards) -1
		Swap the current card and the card at cards[randomNumber]

		func (d deck) shuffle() {
		for index := range d {
			newPosition := rand.Intn(len(d) - 1)

			d[index], d[newPosition] = d[newPosition], d[index]
		}
	}

*/