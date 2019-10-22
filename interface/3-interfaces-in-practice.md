# Interfaces In Practice

```

  type bot interface {
    getGreeting() string
  }

  type frenchBot struct{}
  type englishBot struct{}

  func main() {

    eb := englishBot{}
    fb := frenchBot{}

    printGreeting(eb)
    printGreeting(fb)

  }

  func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
  }

```

To whom it may concern...

**type bot interface**

Our program has a new type called 'bot'.

**getGreeting() string**

If you are a type in this program with a function called 'getGreeting' and you return a string then you are now an honorary member of type 'bot'

Now that you are also an honorary member of type 'bot', you can now call this function called 'printGreeting'

**func printGreeting(b bot)**
