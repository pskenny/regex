package regex

// InfixIntoPostfix takes string with infix notation and returns string with postfix notation.
// Shunting-yard algorithm implementation in video "Shunting yard algorithm in Go" by Ian McLoughlin
func InfixIntoPostfix(infix string) string {
	// Map with runes and associated weighting
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	postfix, stack := []rune{}, []rune{}

	// Throw away index (_), r is rune converted from string
	for _, r := range infix {
		switch {
		case r == '(':
			// Add to end of stack for later use
			stack = append(stack, r)
		case r == ')':
			// "Pop" things off stack until we get to the opening bracket
			for stack[len(stack)-1] != '(' {
				// Add to postfix
				postfix = append(postfix, stack[len(stack)-1])
				// Remove from stack by slicing array
				stack = stack[:len(stack)-1]
			}
			// Preceding for loop stops at open bracket but doesn't remove it, slice array to remove
			stack = stack[:len(stack)-1]
		// Check if rune r is in specials (if requested key doesn't exist it returns 0)
		case specials[r] > 0:
			// Check stack has at least one element and while the precedence of the current rune being read
			// is <= the top element of the stack
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				// Add to postfix
				postfix = append(postfix, stack[len(stack)-1])
				// Remove from stack by slicing array
				stack = stack[:len(stack)-1]
			}
			// Element at top of stack now has less precedence than current rune, add to stack
			stack = append(stack, r)
		default:
			// None operator, add to postfix
			postfix = append(postfix, r)
		}
	}

	// If there's anything left on the stack put it in postfix
	for len(stack) > 0 {
		// Add to postfix
		postfix = append(postfix, stack[len(stack)-1])
		// Remove from stack by slicing array
		stack = stack[:len(stack)-1]
	}

	return string(postfix)
}
