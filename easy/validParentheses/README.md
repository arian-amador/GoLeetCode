# EASY - Valid Parentheses [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Given a string containing just the characters `(`, `)`, `{`, `}`, `[` and `]`, determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Note that an empty string is also considered valid.

##### Example

```Go
Input: "()"
Output: true

Input: "(]"
Output: false

Input: "{[]}"
Output: true
```

### Solution

This is problem requires a stack to store some of our bytes as we iterate over the given string.
One of the key ideas is that we only have to store the open bracket bytes. This is because we will always
want to check for a match if we come across a close bracket byte.

```Go
for i := range s {
	if isOpen(s[i]) {
		stack.push(s[i])
		continue
	} else if isClose(s[i]) {
		// Check if close bracket found at first byte
		if stack.empty() {
			return false
		}
		// check if latest close byte has open matching bracket
		b := stack.pop()
		if !isPairMatch(b, s[i]) {
			return false
		}
	} else {
		return false
	}
}
```

Image a stack that we push open bracket bytes into. Then when we get to a close bracket we `pop` the last open bracket from the top of our stack and validate they are a matching pair.

The rest of the code here is the code to build the stack data structure and some helpful functions to make our code more legible and easier to work with. Functions like `isEmpty()` and `isPairMatch()` are required but make the main search loop a lot easier to read through.

```Go
func isPairMatch(o, c byte) bool {
	if o == '{' && c == '}' {
		return true
	}
	if o == '[' && c == ']' {
		return true
	}
	if o == '(' && c == ')' {
		return true
	}
	return false
}

func isClose(b byte) bool {
	if b == ')' || b == '}' || b == ']' {
		return true
	}
	return false
}

func isOpen(b byte) bool {
	if b == '(' || b == '{' || b == '[' {
		return true
	}
	return false
}
```

The stack data structure code is really straight forward. We have a `struct` that holds our `[]byte` slice. There are some added methods added to the struct for popping and pushing elements to our stack/slice, and some added helpful ones to debug, check the length, and to let us know if the stack is currently `empty()`.

```Go
type stack struct {
	bytes []byte
}

func (s *stack) empty() bool {
	return s.len() == 0
}

func (s *stack) len() int {
	return len(s.bytes)
}

func (s *stack) push(b byte) {
	s.bytes = append(s.bytes, b)
}

func (s *stack) pop() byte {
	i := len(s.bytes) - 1
	b := s.bytes[i]
	s.bytes = s.bytes[:i]
	return b
}

func (s stack) String() string {
	return fmt.Sprintf("length: %d content: %v", s.len(), s.bytes)
}
```

##### Memory Heavy Solution (for fun)

We can use some of Go's strings functions to iterate over the string while removing pairs. At each iteration it checks the length the resulting string, if we have any more pairs, or else we're done.

```Go
func isValid(s string) bool {
	for {
		if len(s) == 0 {
			return true
		}
		if strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
			s = strings.Replace(s, "()", "", -1)
			s = strings.Replace(s, "{}", "", -1)
			s = strings.Replace(s, "[]", "", -1)
		} else {
			return false
		}
	}
}
```

---

#### Hope you find this useful!
