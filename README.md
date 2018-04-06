# Go Regular Expression Library

Project done for Graph Theory, Software Development Year 3, GMIT.

## Example Use

Run the command-line

```
> go get github.com/pskenny/regex
```

Now the library can be used in code

```go
package main

import (
	"fmt"

	paulsregex "github.com/pskenny/regex"
)

func main() {
	// Test infix to postfix
	fmt.Println(paulsregex.InfixIntoPostfix("a.b.c*"))
	fmt.Println(paulsregex.InfixIntoPostfix("(a.(b|d))*"))
}

```