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
	alex := person{firstName: "Alex", lastName: "Anderson"} // Alternative way

	fmt.Println(alex)
}
