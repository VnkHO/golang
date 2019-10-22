# Structs with Receiver Functions

```

  type contactInfo struct {
    email   string
    zipCode int
  }

  type person struct {
    firstName string
    lastName  string
    contactInfo
  }

function main() {

  kim := person{
		firstName: "Kim",
		lastName:  "Gary",
		contactInfo: contactInfo{
			email:   "kim.gary@gmail.com",
			zipCode: 75000,
		},
	}
	kim.updateName("Kimmy")
	fmt.Printf("%+v", kim)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
  fmt.Printf("%+v", p)
}

```

We do not get Kimmy.
We do need pointer.
