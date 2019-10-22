# Pass by value

How RAM / Merory works on the machine (Quick Overview)

**RAM**

| Address | Value |
| ------- | ----- |
| 0000    |       |
| 0001    |       |
| 0002    |       |
| 0003    |       |
| 0004    |       |
| 0005    |       |

Memory on the machine can be thought of as like a bunch of little coveys or a bunch of little slots or a bunch of little boxes.

Each box in the computer's memory can store some data.

And each one of theses little boxes there are these little value containers has some discrete address.

So whenever your program says,
I want to retrieve some information from the computer's memory,
It goes find some address and the it pulls the value out of there.

```

  kim := person{}

```

**RAM**

| Address | Value                       |
| ------- | --------------------------- |
| 0000    |                             |
| 0001    | person{firstName: 'Kim'...} | <--- kim |
| 0002    |                             |
| 0003    |                             |
| 0004    |                             |
| 0005    |                             |

So when we create this new struct of type person
GO will create that struct.
It will then go to the local memory on our own machine and it will attempt to find some container or some spot that is free and has the ability to accept some data.

```

  jim.updateName("Kimmy")

  func (p person) updateName() {}

```

**RAM**

| Address | Value                       |
| ------- | --------------------------- |
| 0000    |                             |
| 0001    | person{firstName: "Kim"...} | <--- kim |
| 0002    |                             |
| 0003    |                             |
| 0004    | person{firstName: "Kim"...} | <---- p |
| 0005    |                             |

GO is what we refer to as a **pass by value** language.
Pass by value means that whenever we pass some value into a function.
GO will take that value or take that struct.

It's going to copy all of that data inside that struct and then place it inside of new some new container inside of our computers memory.

So when we passe _Kim_ into this _updateName function_; Kim still exists by itself as a struct with the firstName of _Kim_ at the address 0001

But GO copy's that value, it finds some other container that is empty and it stuffs that copy into that container and then it runs the code inside of _updateName_ with **this receiver pointing** at that copy.

And so when we modify that field of _firstName_ inside of that function, when we run this code:

```

  func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
  }

```

**We are not updating the original struct of Kim**

We are updating the copy that was just made for our particular function call.
