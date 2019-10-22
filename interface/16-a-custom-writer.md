# A Custom Writer

```

  package main

  import (
    "fmt"
    "io"
    "net/http"
    "os"
  )

  type logWriter struct{}

  func main() {
    response, err := http.Get("https://www.google.fr")
    if err != nil {
      fmt.Println("Error :‡", err)
      os.Exit(1)
    }

    // bs := make([]byte, 99999)
    // response.Body.Read(bs)
    // fmt.Println(string(bs))

    lw := logWriter{}

    io.Copy(lw, response.Body)

  }

  func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Just wrote this many bytes:", len(bs))
    return len(bs), nil
  }


```
