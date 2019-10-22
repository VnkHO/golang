# Gotchas with Pointers

```

func main() {
  mySlice := []string{"Hi", "There", "You"}

  updateSlice(mySlice)
  fmt.Println(mySlice)
}

func updateSlice(s []string) {
  s[0] = "Bye"
}

[Bye There You]

```

This is not working like a struct as with a slice.
