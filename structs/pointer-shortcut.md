# Pointer Shortcut

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

	kim.updateName("Kimmy")
	kim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}


```

With GO, if we define a receiver with a type of pointer to whatever like pointer of blank any type you can possibly imagine.

When we attempt to call this function or we attempt to call his method right here; GO will allow us to either call this function with a pointer or with a chord; like you know the root type person in this case.

```

  kimPointer := &kim
  kimPointer.updateName("Kimmy")

```

kimPointer -> Type of *person, or a pointer to a person

```

  kim.updateName("Kimmy")

```

kim -> Type of person


```

  func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
  }

```

*person -> Type of *person, or a pointer to a person.


