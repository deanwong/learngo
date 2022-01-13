package main

import (
	"fmt"
)

func isValid(s string) bool {
	pairs := map[byte]byte{')': '(', '}': '{', ']': '['}
	n := len(s)
	if n%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		ch := s[i]
		if left, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // false
	fmt.Println(isValid("(]"))     // false
}
