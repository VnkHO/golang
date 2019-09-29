/*

	What is a receiver ?

	func (d deck) print() {}

	Any variable of type "deck" now gets access to the "print" method

	Purpose of a receiver:
		The receiver sets up methods on variables that we create.

		func (cards deck) print() {} cards is like this in JavaScript
		func (d deck) print() {
			for i, card := range d {
				fmt.Println(i, card)
			}
		}

		d -> The actual copy of the deck we are working with is available in the function as a variable called 'd'
			- The variable d right here is a reference to the actual kind of like working variable or the instance of the deck variable that we are working with
		deck -> Every variable of type 'deck' can call this function on itself



*/