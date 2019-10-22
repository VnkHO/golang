# More on the Reader Interface

             Thing that implements Reader

              Read([]byte) (int, err)

Byte Slice ------> Byte Slice <---------- Raw Body of response

Thing that wants to read the body (something that wants to see the Reader interface)
