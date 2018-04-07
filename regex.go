package regex

// Compile takes pattern string and returns NFA
func Compile(pattern string) *NfaFragment {
	return postfixRegexToNfa(infixIntoPostfix(pattern))
}

// Match takes string to check pattern against
func (nfa *NfaFragment) Match(text string) bool {
	match := false
	currentStates := []*state{}
	nextStates := []*state{}

	// Find the initial (current) state(s)
	currentStates = addState(currentStates[:], nfa.initial, nfa.accept)

	// Go through each character in text
	for _, r := range text {
		// Check all current states if rune matches symbol (appropriate for next)
		for _, current := range currentStates {
			// Set up next states
			if current.symbol == r {
				nextStates = addState(nextStates[:], current.edge1, nfa.accept)
			}
		}
		// Swap current and next states, reset next
		currentStates, nextStates = nextStates, []*state{}
	}

	for _, current := range currentStates {
		// Check if any current states are in an accept state
		// Set match to true and break from loop
		if current == nfa.accept {
			match = true
			break
		}
	}

	return match
}

// Regular expression match and addState functions from course material video "Regex match function"

// Recursively traverse states
func addState(states []*state, check *state, accept *state) []*state {
	// Add state to array
	states = append(states, check)

	// Check if it has an empty string pass through (symbol == 0, not in accept)
	if check != accept && check.symbol == 0 {
		// Follow first edge
		states = addState(states, check.edge1, accept)
		// Follow second edge if present
		if check.edge2 != nil {
			states = addState(states, check.edge2, accept)
		}
	}

	return states
}

// Match takes pattern string, text to evaluate, returns in pattern matches in text
func Match(pattern string, text string) bool {
	match := false
	patternNfa := postfixRegexToNfa(infixIntoPostfix(pattern))

	currentStates := []*state{}
	nextStates := []*state{}

	// Find the initial (current) state(s)
	currentStates = addState(currentStates[:], patternNfa.initial, patternNfa.accept)

	// Go through each character in text
	for _, r := range text {
		// Check all current states if rune matches symbol (appropriate for next)
		for _, current := range currentStates {
			// Set up next states
			if current.symbol == r {
				nextStates = addState(nextStates[:], current.edge1, patternNfa.accept)
			}
		}
		// Swap current and next states, reset next
		currentStates, nextStates = nextStates, []*state{}
	}

	for _, current := range currentStates {
		// Check if any current states are in an accept state
		// Set match to true and break from loop
		if current == patternNfa.accept {
			match = true
			break
		}
	}

	return match
}

// Implementation of Thompsons Construction from course material video "go-thompson-final"

type state struct {
	symbol rune
	// Edges equivalent to arrows diagramatically
	edge1 *state
	edge2 *state
}

// NfaFragment connects states together
type NfaFragment struct {
	initial *state
	accept  *state
}

// postfixRegexToNfa takes postfix regular expression string and returns an NFA
func postfixRegexToNfa(postfix string) *NfaFragment {
	nfaStack := []*NfaFragment{}

	// Loop through expression one rune (r) at a time
	for _, r := range postfix {
		switch r {
		// Concatenate
		case '.':
			// "Pop" off the NFA stack (get fragment and slice stack)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			// Join fragments (change pointer in struct)
			fragment1.accept.edge1 = fragment2.initial
			// "Push" joined fragments to NFA stack
			nfaStack = append(nfaStack, &NfaFragment{initial: fragment1.initial, accept: fragment2.accept})
		// Or
		case '|':
			// "Pop" off the NFA stack (get fragment and slice stack)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			// See course material slides.pdf ("Thompson’s construction"), slide 6, for diagramic view of following
			// New initial state
			initial := state{edge1: fragment1.initial, edge2: fragment2.initial}
			accept := state{}
			// Join accept states
			fragment1.accept.edge1 = &accept
			fragment2.accept.edge1 = &accept
			// "Push" new fragment to NFA stack
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial, accept: &accept})
		// Kleene star
		case '*':
			// "Pop" off the NFA stack (get fragment and slice stack)
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			// See course material slides.pdf ("Thompson’s construction"), slide 7, for diagramic view of following
			accept := state{}
			initial := state{edge1: fragment.initial, edge2: &accept}
			fragment.accept.edge1 = fragment.initial
			fragment.accept.edge2 = &accept
			// "Push" new fragment to NFA stack
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial, accept: &accept})
		// Not a special character
		default:
			// "Push" to the stack
			accept := state{}
			// New state with symbol value of rune
			initial := state{symbol: r, edge1: &accept}
			// "Push" to stack
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial, accept: &accept})
		}
	}

	return nfaStack[0]
}

// infixIntoPostfix takes string with infix notation and returns string with postfix notation.
// Shunting-yard algorithm implementation in video "Shunting yard algorithm in Go" by Ian McLoughlin
func infixIntoPostfix(infix string) string {
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
