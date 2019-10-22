# Reading the Docs

```

  Response Struct:

  Status      --> string
  StatusCode  --> int
  Bofy        --> io.ReadCloser

```

```

  io.ReadCloser -->:

  io.ReadCloser Interface

  Reader  -->   io.Reader Interface
                Read([]byte) (int, error)

  Closer  -->   io.Closer Interface
                  Close() (error)

```
