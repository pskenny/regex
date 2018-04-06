# Go Regular Expression Library

Project for Graph Theory, Software Development Year 3, GMIT. Special characters supported are:

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

## How It Works

It is composed of three main parts:

- Infix to postfix expression parser using the [Shunting-yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
- Thompson Construction turns a postfix regular expression into a Non-deterministic Finite Automata (NFA)
- An exported match function traverses the NFA over a given input string and returns a true or false value (match or no match)

## References

- [Official Go Documentation](https://golang.org/doc/)
- [Stack Overflow](https://stackoverflow.com/questions/3639574/writing-a-parser-for-regular-expressions)
- [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
- [Write Your Own Regular Expression Parser](https://www.codeguru.com/cpp/cpp/cpp_mfc/parsing/article.php/c4093/Write-Your-Own-Regular-Expression-Parser.htm)
