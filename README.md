# godc

A dumb "dc" parser.

I read [this nice article](https://eklitzke.org/summing-integer-ranges-with-dc) on summing integer ranges with [dc](https://en.wikipedia.org/wiki/Dc_(computer_program)) and thought it'd be fun to implement an interpreter for the language. Except, lazy as I am I didn't want to read the spec so my implementation is solely based on this article (and I didn't actually test if it'd work with anything else than `36[d1-d1<F+]dsFxp`).

You can try it using `go run main.go 36[d1-d1<F+]dsFxp`.

Also the code is pretty bad, instead of constructing an AST then walking it, the stack is updated as the parser goes through the code (and it's my first ever go code so it doesn't help).
