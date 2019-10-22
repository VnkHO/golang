# Pointer Operations

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

### &variable:

Give me the memory address of the value this variable is pointing at

When we do that, we are saying:
Look at this variable and give us access to the memory address that this variable is pointing at.

```

  kimPointer := &kim

```

**RAM**

|                 | Address | Value                       |          |
| --------------- | ------- | --------------------------- | -------- |
|                 | 0000    |                             |          |
| kimPointer ---> | 0001    | person{firstName: "Kim"...} | <--- Kim |
|                 | 0002    |                             |          |
|                 | 0003    |                             |          |

So in other words if we have &kim; kim is pointing at our struct sitting in memory and that struct is existing at some particular RAM address.

So when we write &kim; we say give me access to the memory address that this struct is sitting at.

example :
0001 takes that value and then we assign it to kimPointer

So now this _kimPointer variable_ is **not pointing** or it is **not a reference** directly to the struct.

Instead it is a **reference to the memory address** that the struct exists at.

If we were to konw print out the value of _kimPointer_; we would probably see some type of memory address (like 0x0ABC123)

---

### \*pointer

Give me the value this memory address is pointing at

So whenever we use the **\*** operator we place the star and then a memory address or a pointer.

Whenever we do that we are saying takes this memory address ang give me the value that exists at that memory address.

```

  func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
  }

```

**(\*pointerToPerson)** is the memory address that _kim_ exists at.

So by saying **\*pointerToPerson**, we are saying here's the pointer I don't want to look at the memory address anymore.
Instead give me direct access to whatever this thing or whatever value is actually sitting here.

So **\*kimPointer** turns into the actual struct of type person

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

Step by step:

```
	kimPointer := &kim

```

_kim_ is a reference to the struct in memory (the actual value of the struct)
_&kim_ we turn it into a memory address or a pointer
We then assign it to _kimPointer_

```

	kimPointer.updateName("Kimmy")

  func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

```

We change the receiver type here to the \* operator.

**\*person** as being a pointer that points at a person.

```

  func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
  }

```

**\*person** -> This is a type description - It means we are working with a **pointer** to a person

**\*pointerToPerson** -> This is an operator - It means we want to manipulate the value the pointer is referencing.

So we have the _updateName function_ which can be called with any receiver of type \*pointer to **\*person**; which is exactly what **kimPointer** is

```

	kimPointer.updateName("Kimmy")

  func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

```

So we call **kimPointer.updateName**; this memory address is then passed into this function as **pointerToPerson**

```

  func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

```

So we can imagine **kimPointer** and **pointerToPerson** are the exact same thing right now.

```

  func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

```

So inside the function, we say **\*pointerToPerson**;

So remember what \* and then an actual pointer means; it says take this memory **\*pointerToPerson** and turn it into an actual value.

So **(\*pointerToPerson)** gets turned into the actual _kim_ person that is sitting in memory.

Some couple of points:

| 0001 | person{firstName: "Kim"...} |
| ---- | --------------------------- |


0001 -> Address

person{firstName: "Kim"...} -> value

Turn **address** into _value_ with **\*address**

**\*address** Turn **address** into _value_

---

Turn _value_ into **address** with _&value_

_&value_ --> Turn _value_ into **address**
