# Go Regular Expression Library

Project for Graph Theory, Software Development Year 3, GMIT. Special characters in the supported are:

- "." for concatenation
- "|" for or
- "*" for [Kleene star](https://en.wikipedia.org/wiki/Kleene_star)

It was written with use of course material provided and references (see bottom).

## Example Use

Ensure Go language is [properly installed](https://golang.org/doc/install). Run in the command-line

```shell
> go get github.com/pskenny/regex
```

Now the library can be used in code, as in the following example:

```go
package main

import (
    "fmt"

    paulsregex "github.com/pskenny/regex"
)

func main() {
    fmt.Println(paulsregex.Match("a.b*", "abbb"))
}

```