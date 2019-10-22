# Rules of Interfaces

```

  type bot interface {
    getGreeting(string, int) (string, error)
  }

```

**bot** -> Interface name
**getGreeting** -> Function name
**string, int** -> List of argument types
**string, error** -> List of return types

```

  type bot interface {
    getGreeting(string, int) (string, error)
    getBotVersion() float64
  }

```

| Concrete Type | Interface Type |
| ------------- | -------------- |
| map           |                |
| struct        | bot            |
| int           |                |
| string        |                |
| englishBot    |                |
| frenchBot     |                |

A concrete type is something that we can actually kind of create a value out of directly and then access it and change it and create new copies of it and whatnot

Interface Type, we can't actually create a value directly out of this type.
