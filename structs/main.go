package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Defining a new custom type
type person struct {
	firstName string
	lastName  string
	contactInfo
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

	kim := person{
		firstName: "Kim",
		lastName:  "Gary",
		contactInfo: contactInfo{
			email:   "kim.gary@gmail.com",
			zipCode: 75000,
		},
	}
	kim.updateName("Kimmy")
	kim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*

	Structs
		Data structure. Collection of propertiees that are related together.

*/

// You can think a struct is "like" an object in JavaScript
