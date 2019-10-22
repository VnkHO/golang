# Reference vs Value Types

| Array                     | Slices                                     |
| ------------------------- | ------------------------------------------ |
| Primitive Data structures | Can grow and shrink                        |
| Can't be resized          | Used 99% of the time for lists of elements |
| Rarely used directly      |                                            |

So when we create a slice; GO internally is creating **two seperate data structures** for us.

```

  mySlice := []string{"Hi", "There", "You"}

```

The first is what we refer to as the slice;
The slice is a data structure that has three elements inside:

- pointer to head
  - A pointer over to the underlying array that represents the actual list of items.
- capacity
  - How many elements it can contain at present.
- lenght
  - Length represents how many elements currently exist inside the slice.

**array**

"Hi" | "There" | "You"

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

**Reference Types** -> Don't worry about pointers

**Value Types** -> Use pointers to change theses things in a function
