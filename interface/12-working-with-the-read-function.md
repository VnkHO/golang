# Working with the Read Function

```

	bs := make([]byte, 99999)

```

This **make** function is a built in function in the language that takes a type of a slice.
And then as a second argument this is the number of elements or empty spaces that we want to slice to be initialized with.

```

	bs := make([]byte, 99999)
	response.Body.Read(bs)
	fmt.Println(string(bs))


```
