# Structs with Pointers

GO is a **pass by value** language that means that whenever we pass a value into a function; GO copies that value and then the copy is made available to the code inside the function.

How to fix this:

```

func main() {

	kim := person{
		firstName: "Kim",
		lastName:  "Gary",
		contactInfo: contactInfo{
			email:   "kim.gary@gmail.com",
			zipCode: 75000,
		},
	}
	kimPointer := &kim
	kimPointer.updateName("Kimmy")
	kim.print()
}

  func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
  }


```
