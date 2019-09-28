/*

	func main() {
		cards := newCard()

		fmt.Println(card)
	}

	func newCard() string {
		return "Five of Diamonds"
	}

	The error that we saw said that the function that we rode out here
	was attempting to return a value of type string when we had said that nothing would be returned

	In order to tell Golang a lliter bit more about the function we just made
	It requires us to always be very EXPLICIT and TELL IT exactly what DATA TYPE we are going to return from any giving function

		func newCard() string {}

		newCard -> Define a function called 'newCard'
		string	-> When executed, this function will return a value of type 'string'

*/