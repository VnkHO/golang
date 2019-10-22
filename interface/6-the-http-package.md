# The HTTP package

```

  package main

  import (
    "fmt"
    "net/http"
    "os"
  )

  func main() {
    response, err := http.Get("https://www.google.fr")
    if err != nil {
      fmt.Println("Error :", err)
      os.Exit(1)
    }
    fmt.Println(response)
  }


```
