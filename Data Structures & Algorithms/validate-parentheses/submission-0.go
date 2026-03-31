func isValid(s string) bool {
	stack := []rune{}
	lookup := map[rune]rune{
      '(': ')',
      '{': '}',
      '[': ']',
  	}

	for _,r := range s {
		// check if opening char
		if _,ok := lookup[r]; ok {
			// opening char, add to stack
			stack = append(stack, r)
		} else {
			if len(stack) == 0 { return false }
			// closing char
			// check if char matches top of stack
			if lookup[stack[len(stack)-1]] == r {
				stack = stack[:len(stack)-1]
			} else {
				// invalid char
				return false
			}
		}
	}
	return len(stack) == 0	 
}


// Push: Use append() to add elements to the end of the slice. 
// Pop: Slice off the last element using slice[:len(slice)-1]. 
// Peek: Access the last element via slice[len(slice)-1].