# The Writer Interface

```

	io.Copy(os.Stdout, response.Body)


```

Source of data --> Reader --> []byte (Output data that anyone can work with)

[]byte --> Writer --> Some form of output

---

[] byte -> Writer -> Source of Output

Writer: Writer interface describes something that can take info and send it outside of our program

Source Of Output: We need to find something in the standard library that \*implements\* the Writer interface, and use that to log out all the data that we are receiving from the Reader
