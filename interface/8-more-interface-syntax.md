# More Interface Syntax

If we specify a interface as a value inside of a struct, all we are doing is we are saying that the body (Body) field can have any value assigned to it so long as it fulfills this interface (io.readCloser).

In GO, we can take multiple interfaces so different interfaces and assemble them together to form another interface.

So both **Reader** and **Closer**

The **ReadCloser** interface says if you want to fulfill the **ReadCloser** interface like if you want to satisfy the requirements of this interface, you have to satisfy the requirements of both the **Reader** and the **Closer**

So all we are really doing here is defining a new interface by putting together pieces of other ones inside of our application.

So we can freely embed one interface into another as much as we please as long as it acutally serves the purpose of building our application.

What really matters to us, is what the **Reader** interface and **Closer** itnerface are requiring of us.
