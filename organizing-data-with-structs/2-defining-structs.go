/*

	Whenever we make a struct we have to first define all of the different properties that a struct might have.package structs

	We're going to provide this in some type of rule set to go and then we can wreate a value that matches that type of structure definition.


	First thing we are going to do is we are going to tell Go exactly what different filds a person should have

		Tell Go what fields the person struct has

			firstName -> <string>
			lastName	-> <string>
					person struct


		Step 2:

			Create a new value of type person

				firstName -> "Alex"
				lastName	-> "Anderson"

	// Defining a new custom type
	type person struct {
		firstName string
		lastName  string
	}

*/