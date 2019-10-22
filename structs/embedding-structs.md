# Embedding structs

```

  type contactInfo struct {
    email   string
    zipCode int
  }

  type person struct {
    firstName string
    lastName  string
    contact   contactInfo
  }

  kim := person{
		firstName: "Kim",
		lastName:  "Gary",
		contact: contactInfo{
			email:   "kim.gary@gmail.com",
			zipCode: 75000,
		},
	}
	fmt.Printf("%+v", kim)

```

The ability eo embed one struct inside of another. 