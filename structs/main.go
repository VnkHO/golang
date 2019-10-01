package main

import "fmt"

// Defining a new custom type
type person struct {
	firstName string
	lastName  string
}

func main() {

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
}
