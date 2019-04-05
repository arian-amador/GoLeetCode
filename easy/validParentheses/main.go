package main

import (
	"fmt"
)

func main() {}

func isValid(s string) bool {
	// Empty strings are valid
	if len(s) == 0 {
		return true
	}

	stack := stack{}

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

	return stack.empty()
}

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
