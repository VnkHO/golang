/*

	When you create a variable in Go and you do not actually assign any value to it
	or you don't actually know kind of pre-populated these different fields
	Go assigns what is referred to as a "zero value" to each of these different fields inside the struct

	Type				Zero Value

	String 	->	""
	int 		->	0
	float		->	0
	bool		->	0


	// order of definition of fields
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"} // Alternative way

	// Third way
	var alex person

	// Updating
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)


*/