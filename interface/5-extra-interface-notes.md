# Extra Interface Notes

### Interfaces are **not** generic types

Other languages have 'generic' types - GO (famously) does not.

### Interfaces are 'implicit'

We don't manually have to say that our custom type satisfies some interface.

### Interfaces are a contract to help us manage types

GARBAGE IN -> GARBAGE OUT.
If our custom type's implementation of a function is broken then interfaces won't help us!

### Interfaces are tough. Step #1 is understanding how to read them.

Understand how to read interfaces in the standard lib. Writing your own interfaces is tough and requires experience.
