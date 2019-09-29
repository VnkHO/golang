/*

	Array -> Fixed length list of things
	Slice -> An array that can grow or shrink

	Slices and Arrays both lust be defined with a Data Type

	Slices:

		- Every element in a slice must be of same type

		example:
			cards := []string{"Ace of Diamonds", newCard()}

		How to add an element to slice
			- cards = append(cards, "Six of Spades");
		So this will take the cards slice it will add on a new string to the very end
		Append function doest not modify the existing slice
		returns a new slice that we then assign back to a variable of cards


			for i, card := range cards {
				fmt.Println(i, card)
			}

			index -> index of this element in the array
			card -> current card we are iteration over
			range cards -> Take the slice of 'cards' and loop over it

			fmt.Println(i, card) -> Run this one time for each card in the slice

			:= with for loops, every single time that we step through this list of cards
			we are really throwing away the previous index and the previous card that had been declared


			Array:
				The type [n]T is an array of n values of type T.
			The expression: var a [10]int
				declares a variable a as an array of ten integers.

			An array's length is part of its type, so arrays cannot be resized.

			Example:
				func main() {
					var a [2]string
					a[0] = "Hello"
					a[1] = "World"
					fmt.Println(a[0], a[1])
					fmt.Println(a)

					primes := [6]int{2, 3, 5, 7, 11, 13}
					fmt.Println(primes)
				}


*/